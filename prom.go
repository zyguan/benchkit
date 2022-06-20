package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

type PromMetricLine struct {
	Metric map[string]string `json:"metric"`
	Values [][2]any          `json:"values"`
	Rebase uint64            `json:"-"`
}

func (m PromMetricLine) Empty() bool {
	for _, tuple := range m.Values {
		v, err := strconv.ParseFloat(tuple[1].(string), 64)
		if err == nil && !math.IsNaN(v) && !math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func (m PromMetricLine) MarshalJSON() ([]byte, error) {
	ts, vs := make([]uint64, 0, len(m.Values)), make([]float64, 0, len(m.Values))
	var startTime uint64
	for i, tuple := range m.Values {
		t, ok := tuple[0].(float64)
		if !ok {
			continue
		}
		v, err := strconv.ParseFloat(tuple[1].(string), 64)
		if err != nil || math.IsNaN(v) || math.IsInf(v, 0) {
			continue
		}
		epochMillis := uint64(t * 1000)
		if i == 0 {
			startTime = epochMillis
		}
		if m.Rebase > 0 {
			epochMillis = m.Rebase*1000 + epochMillis - startTime
		}
		ts = append(ts, epochMillis)
		vs = append(vs, v)
	}
	var obj struct {
		Metric     map[string]string `json:"metric"`
		Timestamps []uint64          `json:"timestamps"`
		Values     []float64         `json:"values"`
	}
	obj.Metric = m.Metric
	obj.Timestamps = ts
	obj.Values = vs
	return json.Marshal(obj)
}

func PromListMetrics(endpoint string, headers map[string]string) ([]string, error) {
	req, err := promNewGetRequest(endpoint, headers, "/api/v1/label/__name__/values")
	if err != nil {
		return nil, err
	}
	data, err := promDoRequest[[]string](req)
	if err != nil {
		return nil, err
	}
	return *data, nil
}

func PromListSeries(endpoint string, headers map[string]string, start int64, end int64, match string, matches ...string) ([]map[string]string, error) {
	queries := make([]string, 0, len(matches)+3)
	if start > 0 {
		queries = append(queries, "start", strconv.FormatInt(start, 10))
	}
	if end > 0 {
		queries = append(queries, "end", strconv.FormatInt(end, 10))
	}
	queries = append(queries, "match[]", match)
	for _, m := range matches {
		queries = append(queries, "match[]", m)
	}
	req, err := promNewGetRequest(endpoint, headers, "/api/v1/series", queries...)
	if err != nil {
		return nil, err
	}
	data, err := promDoRequest[[]map[string]string](req)
	if err != nil {
		return nil, err
	}
	return *data, nil
}

func PromQueryMatrix(endpoint string, headers map[string]string, query string, start int64, end int64, step int64) ([]PromMetricLine, error) {
	if step == 0 {
		step = 15
	}
	if end == 0 {
		end = time.Now().Unix()
	}
	if start == 0 {
		start = end - 3600
	}
	req, err := promNewGetRequest(endpoint, headers, "/api/v1/query_range",
		"query", query,
		"start", strconv.FormatInt(start, 10),
		"end", strconv.FormatInt(end, 10),
		"step", strconv.FormatInt(step, 10),
	)
	if err != nil {
		return nil, err
	}
	data, err := promDoRequest[struct {
		ResultType string           `json:"resultType"`
		Result     []PromMetricLine `json:"result"`
	}](req)
	if err != nil {
		return nil, err
	}
	if data.ResultType != "matrix" {
		return nil, fmt.Errorf("unexpected result type: %q", data.ResultType)
	}
	return data.Result, nil
}

func promNewGetRequest(endpoint string, headers map[string]string, path string, queries ...string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint+path, nil)
	if err != nil {
		return nil, err
	}
	if len(queries) > 0 {
		q := req.URL.Query()
		for i := 0; i+1 < len(queries); i += 2 {
			q.Add(queries[i], queries[i+1])
		}
		req.URL.RawQuery = q.Encode()
	}
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	return req, nil
}

func promDoRequest[T any](req *http.Request) (*T, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body struct {
		Status string `json:"status"`
		Data   T      `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	if body.Status != "success" {
		return nil, fmt.Errorf("unexpected response status: %q", body.Status)
	}
	return &body.Data, nil
}
