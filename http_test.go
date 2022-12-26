package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Yihaaa")
}

func TestHttp(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port+"/hello", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	errHandler(err)
	bodyString := string(body)

	fmt.Println(bodyString)
}
