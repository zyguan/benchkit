package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	_ "embed"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type BenchResult struct {
	ID       string   `json:"id,omitempty"`
	Name     string   `json:"name"`
	Tags     []string `json:"tags"`
	Cmd      []string `json:"cmd"`
	Started  int64    `json:"started"`
	Finished int64    `json:"finished"`
	Exit     int      `json:"exit"`
	Error    string   `json:"error,omitempty"`
	Stdout   string   `json:"stdout,omitempty"`
	Stderr   string   `json:"stderr,omitempty"`
}

func (r BenchResult) ReportTo(endpoint string) error {
	raw, err := json.Marshal(r)
	if err != nil {
		return err
	}
	resp, err := http.Post(endpoint+"/results", "application/json; charset=utf-8", bytes.NewReader(raw))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("unexpected response status(%s): %s", resp.Status, string(body))
	}
	return nil
}

func BenchRun(ctx context.Context, name string, tags []string, proc string, args ...string) BenchResult {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	result := BenchResult{
		ID:   uuid.NewString(),
		Name: name,
		Tags: tags,
		Cmd:  append([]string{proc}, args...),
	}
	// stderr is the logging output, thus redirect bench stderr to stdout.
	p := exec.CommandContext(ctx, proc, args...)
	p.Stdout = io.MultiWriter(stdout, os.Stdout)
	p.Stderr = io.MultiWriter(stderr, os.Stdout)
	result.Started = time.Now().Unix()
	err := p.Run()
	result.Finished = time.Now().Unix()
	result.Stdout = stdout.String()
	result.Stderr = stderr.String()
	if err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			result.Exit = e.ExitCode()
		} else {
			result.Exit = -1
		}
		result.Error = err.Error()
	}
	return result
}

var (
	//go:embed db/bench_result.sql
	ddlBenchResult string
	//go:embed db/bench_result_tag.sql
	ddlBenchResultTag string
)

func BenchServerInitDB(db *sql.DB) error {
	for _, ddl := range []string{ddlBenchResult, ddlBenchResultTag} {
		_, err := db.Exec(ddl)
		if err != nil {
			return err
		}
	}
	return nil
}

func BenchServerHandler(db *sql.DB) http.Handler {
	store := &benchResultStore{db: db}
	r := mux.NewRouter()
	r.Use(withAccessLog, allowCORS)
	r.HandleFunc("/results", store.handlePutBenchResult).Methods(http.MethodPut, http.MethodPost)
	r.HandleFunc("/results", store.handleListBenchResult).Methods(http.MethodGet)
	return r
}

type benchResultStore struct {
	db *sql.DB
}

func (store *benchResultStore) handleListBenchResult(w http.ResponseWriter, r *http.Request) {
	var (
		query  = strings.TrimSpace(r.URL.Query().Get("query"))
		limit  = 10
		offset = 0

		includeOutput = false
	)
	if r.URL.Query().Has("limit") {
		v, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
		if err != nil {
			httpReturnError(w, http.StatusBadRequest, err)
			return
		}
		limit = int(v)
	}
	if r.URL.Query().Has("offset") {
		v, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
		if err != nil {
			httpReturnError(w, http.StatusBadRequest, err)
			return
		}
		limit = int(v)
	}
	if r.URL.Query().Has("includeOutput") {
		v := r.URL.Query().Get("includeOutput")
		if v == "1" || strings.ToLower(v) == "true" {
			includeOutput = true
		}
	}
	q, args, err := BuildQuerySQL(query, includeOutput, limit, offset)
	if err != nil {
		glog.Warningf("failed to build query from %+v: %v", r.URL.Query(), err)
		httpReturnError(w, http.StatusBadRequest, err)
		return
	}
	glog.Infof("compile query from %+v to %q %v", r.URL.Query(), q, args)
	rs, err := QueryBenchResults(context.Background(), store.db, q, args)
	if err != nil {
		glog.Warningf("failed to execute query: %v", err)
		myerr, ok := err.(*mysql.MySQLError)
		status := http.StatusInternalServerError
		if ok && myerr.Number == 1064 {
			status = http.StatusBadRequest
		}
		httpReturnError(w, status, err)
		return
	}
	httpReturnJSON(w, http.StatusOK, rs)
}

func (store *benchResultStore) handlePutBenchResult(w http.ResponseWriter, r *http.Request) {
	var result BenchResult
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		httpReturnError(w, http.StatusBadRequest, err)
		return
	}
	if len(result.ID) == 0 {
		result.ID = uuid.NewString()
	}
	cmd, _ := json.Marshal(result.Cmd)

	tx, err := store.db.Begin()
	if err != nil {
		httpReturnError(w, http.StatusInternalServerError, err)
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec("insert into bench_result (`id`, `name`, `cmd`, `started`, `finished`, `exit`, `error`, `stdout`, `stderr`) values (?, ?, ?, ?, ?, ?, ?, ?, ?) "+
		"on duplicate key update `name`=values(`name`), `cmd`=values(`cmd`), `started`=values(`started`), `finished`=values(`finished`), `exit`=values(`exit`), `error`=values(`error`), `stdout`=values(`stdout`), `stderr`=values(`stderr`)",
		result.ID, result.Name, string(cmd), result.Started, result.Finished, result.Exit, result.Error, result.Stdout, result.Stderr)
	if err != nil {
		httpReturnError(w, http.StatusInternalServerError, err)
		return
	}
	_, err = tx.Exec("delete from `bench_result_tag` where `id` = ?", result.ID)
	if err != nil {
		httpReturnError(w, http.StatusInternalServerError, err)
		return
	}
	sort.Strings(result.Tags)
	k := 0
	for i, tag := range result.Tags {
		if i > 0 && result.Tags[i-1] == tag {
			continue
		} else {
			result.Tags[k] = tag
			k++
		}
		_, err = tx.Exec("insert into `bench_result_tag` (`id`, `tag`) values (?, ?)", result.ID, tag)
		if err != nil {
			httpReturnError(w, http.StatusInternalServerError, err)
			return
		}
	}
	result.Tags = result.Tags[:k]

	if err = tx.Commit(); err != nil {
		httpReturnError(w, http.StatusInternalServerError, err)
		return
	}
	httpReturnJSON(w, http.StatusOK, result)
}
