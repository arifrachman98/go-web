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
