package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang/glog"
)

var client *http.Client = http.DefaultClient

func withAccessLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			glog.Infof("%s %s [%.2fms]", r.Method, r.RequestURI, float64(time.Since(start))/float64(time.Millisecond))
		}()
		h.ServeHTTP(w, r)
	})
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "*")
		h.ServeHTTP(w, r)
	})
}

type ErrorResponse struct {
	Code  int    `json:"code,omitempty"`
	Error string `json:"error"`
}

func httpReturnJSON[T any](w http.ResponseWriter, status int, body T) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func httpReturnError(w http.ResponseWriter, status int, err error) {
	httpReturnJSON(w, status, ErrorResponse{Error: err.Error()})
}
