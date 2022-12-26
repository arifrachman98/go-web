package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Test-Powered-By", "Arif Rachman Hakim")
	fmt.Fprint(w, "OK")
}

func TestRequestHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port+"/", nil)
	req.Header.Add("Content-type", "application/json")
	rec := httptest.NewRecorder()

	RequestHeader(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	errHandler(err)

	fmt.Println(string(body))
}

func TestResponseHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:"+port+"/", nil)
	req.Header.Add("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	ResponseHeader(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	errHandler(err)

	fmt.Println(string(body))

	fmt.Println(res.Header.Get("Test-Powered-By"))
}
