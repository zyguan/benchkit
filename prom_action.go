package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/golang/glog"
)

type ProgressListener interface {
	Begin(n int)
	Step(i int, text string)
	End()
}

type NoopProgressListener struct{}

func (l NoopProgressListener) Begin(n int)             {}
func (l NoopProgressListener) Step(i int, text string) {}
func (l NoopProgressListener) End()                    {}

type ConsoleProgressListener struct{ n int }

func (l *ConsoleProgressListener) Begin(n int) {
	l.n = n
	fmt.Fprintf(os.Stdout, "[0/%d] starting\n", l.n)
}

func (l *ConsoleProgressListener) Step(i int, text string) {
	fmt.Fprintf(os.Stdout, "\033[A\033[K[%d/%d] %s\n", i, l.n, text)
}

func (l *ConsoleProgressListener) End() {
	fmt.Fprintf(os.Stdout, "\033[A\033[K[%d/%d] done\n", l.n, l.n)
}

type MetricsWriter interface {
	io.Closer
	Write(line PromMetricLine) error
}

type MetricsStore interface {
	Open(name string) (MetricsWriter, error)
}

type metricsFileWriter struct {
	file *os.File
	buf  *bufio.Writer
	zip  *gzip.Writer
	enc  *json.Encoder

	rebase uint64
}

func (f *metricsFileWriter) Write(line PromMetricLine) error {
	line.Rebase = f.rebase
	return f.enc.Encode(line)
}

func (f *metricsFileWriter) Close() error {
	if err := f.zip.Close(); err != nil {
		f.buf.Flush()
		f.file.Close()
		return err
	}
	if err := f.buf.Flush(); err != nil {
		f.file.Close()
		return err
	}
	return f.file.Close()
}

type FileMetricsStore struct {
	dir string
	gz  int
	pb  ProgressListener

	rebase uint64
}

func NewFileMetricsStore(dir string, gz int, rebase uint64) (*FileMetricsStore, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return &FileMetricsStore{dir: dir, gz: gz, pb: NoopProgressListener{}, rebase: rebase}, nil
}

func (store *FileMetricsStore) Open(name string) (MetricsWriter, error) {
	f, err := os.OpenFile(filepath.Join(store.dir, name+".jsonl.gz"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewWriter(f)
	zip := gzip.NewWriter(buf)
	return &metricsFileWriter{file: f, buf: buf, zip: zip, enc: json.NewEncoder(zip), rebase: store.rebase}, nil
}

func (store *FileMetricsStore) ShowProgress(ok bool) {
	if ok {
		store.pb = &ConsoleProgressListener{}
	} else {
		store.pb = NoopProgressListener{}
	}
}

func (store *FileMetricsStore) UploadToVM(endpoint string, headers map[string]string, labels map[string]string, concurrency int) error {
	var (
		wg sync.WaitGroup
		in = make(chan string)
	)
	if concurrency < 1 {
		concurrency = 1
	}
	// validate the endpoint
	if _, err := http.NewRequest(http.MethodPost, endpoint+"/api/v1/import", nil); err != nil {
		return err
	}

	if meta, err := os.ReadFile(filepath.Join(store.dir, "bench.json")); err == nil {
		var info struct {
			ID string `json:"id"`
		}
		err = json.Unmarshal(meta, &info)
		if err == nil && len(info.ID) > 0 {
			if labels == nil {
				labels = make(map[string]string)
			}
			labels["bench"] = info.ID
		}
	}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range in {
				f, err := store.openMetricsFile(n)
				if err != nil {
					glog.Warningf("open %q: %v", n, err)
					continue
				}
				req, err := http.NewRequest(http.MethodPost, endpoint+"/api/v1/import", f)
				if err != nil {
					glog.Warningf("new import request(%q,%q): %v", endpoint+"/api/v1/import", n, err)
					f.Close()
					continue
				}
				req.Header.Set("Content-Encoding", "gzip")
				if len(labels) > 0 {
					vs := req.URL.Query()
					for k, v := range labels {
						vs.Add("extra_label", k+"="+v)
					}
					req.URL.RawQuery = vs.Encode()
				}
				if len(headers) > 0 {
					for k, v := range headers {
						req.Header.Set(k, v)
					}
				}
				if _, err = client.Do(req); err != nil {
					glog.Errorf("do import request(%q): %v", n, err)
				}
				f.Close()
			}
		}()
	}
	ns, err := filepath.Glob(filepath.Join(store.dir, "*.jsonl.gz"))
	if err != nil {
		return err
	}
	store.pb.Begin(len(ns))
	for i, n := range ns {
		in <- n
		store.pb.Step(i+1, filepath.Base(n))
	}
	close(in)
	wg.Wait()
	store.pb.End()
	return nil
}

func (store *FileMetricsStore) openMetricsFile(name string) (io.ReadCloser, error) {
	f, err := os.Open(name)
	if store.rebase == 0 {
		return f, err
	}
	return newRebasedMetricsFile(f, store.rebase)

}

func newRebasedMetricsFile(f *os.File, rebase uint64) (*rebasedMetricsFile, error) {
	r, err := gzip.NewReader(f)
	if err != nil {
		f.Close()
		return nil, err
	}
	buf := new(bytes.Buffer)
	zip := gzip.NewWriter(buf)

	return &rebasedMetricsFile{
		src:    f,
		dec:    json.NewDecoder(r),
		enc:    json.NewEncoder(zip),
		buf:    buf,
		zip:    zip,
		rebase: rebase,
	}, nil
}

type rebasedMetricsFile struct {
	src *os.File
	dec *json.Decoder
	enc *json.Encoder
	buf *bytes.Buffer
	zip *gzip.Writer

	rebase uint64
}

func (f *rebasedMetricsFile) next() bool { return f.dec != nil }

func (f *rebasedMetricsFile) done() {
	f.zip.Close()
	f.dec = nil
}

func (f *rebasedMetricsFile) Read(p []byte) (n int, err error) {
	if f.buf.Len() == 0 && f.next() {
		var line PromMetricLine
		if err := f.dec.Decode(&line); err != nil {
			f.done()
			if errors.Is(err, io.EOF) {
				return f.buf.Read(p)
			}
			return 0, err
		}
		line.Rebase = f.rebase
		if err := f.enc.Encode(line); err != nil {
			f.done()
			return 0, err
		}
	}
	return f.buf.Read(p)
}

func (f *rebasedMetricsFile) Close() error {
	return f.src.Close()
}

type vmWriter struct {
	store  *VictoriaMetricsStore
	lines  []PromMetricLine
	buf    bytes.Buffer
	rebase uint64
}

func (w *vmWriter) Write(line PromMetricLine) error {
	line.Rebase = w.rebase
	w.lines = append(w.lines, line)
	if len(w.lines) >= w.store.Batch {
		return w.Flush()
	}
	return nil
}

func (w *vmWriter) Close() error {
	return w.Flush()
}

func (w *vmWriter) Flush() error {
	if len(w.lines) == 0 {
		return nil
	}
	defer func() { w.lines = w.lines[:0] }()
	w.buf.Reset()
	zip := gzip.NewWriter(&w.buf)
	enc := json.NewEncoder(zip)
	for _, line := range w.lines {
		if err := enc.Encode(line); err != nil {
			return err
		}
	}
	if err := zip.Close(); err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, w.store.Endpoint+"/api/v1/import", &w.buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Encoding", "gzip")
	if len(w.store.Labels) > 0 {
		vs := req.URL.Query()
		for k, v := range w.store.Labels {
			vs.Add("extra_label", k+"="+v)
		}
		req.URL.RawQuery = vs.Encode()
	}
	if len(w.store.Headers) > 0 {
		for k, v := range w.store.Headers {
			req.Header.Set(k, v)
		}
	}
	_, err = client.Do(req)
	return err
}

type VictoriaMetricsStore struct {
	Endpoint string
	Headers  map[string]string
	Labels   map[string]string
	Batch    int
	Rebase   uint64
}

func (store *VictoriaMetricsStore) SetLabel(name string, value string) {
	if store.Labels == nil {
		store.Labels = make(map[string]string)
	}
	store.Labels[name] = value
}

func (store *VictoriaMetricsStore) Open(name string) (MetricsWriter, error) {
	return &vmWriter{store: store, lines: make([]PromMetricLine, 0, store.Batch), rebase: store.Rebase}, nil
}

type PromDumpOptions struct {
	Endpoint string
	Start    int64
	End      int64
	Step     int64
	Selector map[string]string
	Headers  map[string]string

	MaxSteps    int
	MaxSamples  int
	Concurrency int
	Filter      *regexp.Regexp
	Progress    ProgressListener
}

func (opts PromDumpOptions) SetDefaults() PromDumpOptions {
	if opts.End == 0 {
		opts.End = time.Now().Unix()
	}
	if opts.Start == 0 {
		opts.Start = opts.End - 3600
	}
	if opts.Step == 0 {
		opts.Step = 15
	}
	if opts.MaxSteps == 0 {
		opts.MaxSteps = 720
	}
	if opts.MaxSamples == 0 {
		opts.MaxSamples = 50000000
	}
	if opts.Concurrency < 1 {
		opts.Concurrency = runtime.NumCPU() * 2
	}
	return opts
}

func PromDumpMetrics(store MetricsStore, opts PromDumpOptions) error {
	metrics, err := PromListMetrics(opts.Endpoint, opts.Headers)
	if err != nil {
		return err
	}
	var (
		wg sync.WaitGroup
		in = make(chan string)

		status = new(atomic.Value)
		done   = make(chan struct{})
	)
	for i := 0; i < opts.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			selector := make(map[string]string, len(opts.Selector)+1)
			for k, v := range opts.Selector {
				selector[k] = v
			}
			for m := range in {
				selector["__name__"] = m
				w, err := store.Open(m)
				if err != nil {
					glog.Errorf("open %q: %v", m, err)
					continue
				}
				err = promDumpMetric(w, selector, opts, status)
				if err != nil {
					glog.Errorf("dump %q: %v", m, err)
				}
				w.Close()
			}
		}()
	}
	if opts.Filter != nil {
		n := 0
		for _, m := range metrics {
			if opts.Filter.MatchString(m) {
				metrics[n] = m
				n++
			}
		}
		metrics = metrics[:n]
	}
	pb := opts.Progress
	if pb == nil {
		pb = NoopProgressListener{}
	}
	pb.Begin(len(metrics))
	for i, m := range metrics {
		in <- m
		pb.Step(i+1, "dumping "+m)
	}
	close(in)
	go func() {
		wg.Wait()
		close(done)
	}()
wait:
	for {
		select {
		case <-time.After(time.Second):
			if msg, ok := status.Load().(string); ok {
				pb.Step(len(metrics), "waiting "+msg)
			}
		case <-done:
			break wait
		}
	}
	pb.End()
	return nil
}

func promDumpMetric(out MetricsWriter, selector map[string]string, opts PromDumpOptions, status *atomic.Value) error {
	m := promFormatSelector(selector)
	xs, err := PromListSeries(opts.Endpoint, opts.Headers, opts.Start, opts.End, m)
	if err != nil {
		return err
	}
	if len(xs) == 0 {
		return nil
	}

	var queries []string

	if len(xs)*opts.MaxSteps < opts.MaxSamples {
		queries = []string{m}
	} else {
		// try to reduce queries (batch by instance)
		queries = make([]string, len(xs))
		for i, x := range xs {
			queries[i] = promFormatSelector(x, "instance")
		}
		if len(queries) > 1 {
			sort.Strings(queries)
			k := 1
			for i := 1; i < len(queries); i++ {
				if queries[k-1] == queries[i] {
					continue
				}
				queries[k] = queries[i]
				k++
			}
			queries = queries[:k]
		}
		if len(xs)*opts.MaxSteps/len(queries) >= opts.MaxSamples {
			// still too many samples per query
			queries = queries[:0]
			for _, x := range xs {
				queries = append(queries, promFormatSelector(x))
			}
		}
	}

	for i, query := range queries {
		for start := opts.Start; start < opts.End; start += int64(opts.MaxSteps) * opts.Step {
			end := start + int64(opts.MaxSteps-1)*opts.Step
			if end > opts.End {
				end = opts.End
			}
			lines, err := PromQueryMatrix(opts.Endpoint, opts.Headers, query, start, end, opts.Step)
			if err != nil {
				glog.Errorf("query metrics line(%q,%d,%d,%d): %v", query, start, end, opts.Step, err)
				continue
			}
			for _, line := range lines {
				if !line.Empty() {
					if err = out.Write(line); err != nil {
						glog.Errorf("write metrics line(%q,%d,%d,%d): %v", query, start, end, opts.Step, err)
					}
				}
			}
		}
		status.Store(fmt.Sprintf("%s (%d/%d)", selector["__name__"], i+1, len(queries)))
	}
	return nil
}

func promFormatSelector(selector map[string]string, withoutLabels ...string) string {
	var buf strings.Builder
	buf.WriteString(selector["__name__"])
	if len(selector) <= 1 && buf.Len() > 0 {
		return buf.String()
	}
	ks := make([]string, 0, len(selector))
loop:
	for k := range selector {
		if k == "__name__" {
			continue
		}
		for _, l := range withoutLabels {
			if k == l {
				continue loop
			}
		}
		ks = append(ks, k)
	}
	sort.Strings(ks)
	buf.WriteByte('{')
	for i, k := range ks {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(strconv.Quote(selector[k]))
	}
	buf.WriteByte('}')
	return buf.String()
}
