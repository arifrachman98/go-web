package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	errHandler(err)

	firstname := r.PostForm.Get("first_name")
	lastname := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstname, lastname)
}

func TestFormPost(t *testing.T) {
	reqBody := strings.NewReader("first_name=Arif&last_name=Rachman")
	req := httptest.NewRequest(http.MethodPost, "http://localhost:"+port, reqBody)
	req.Header.Add("Content-Type", "application/test-www-form-urlencoded")

	rec := httptest.NewRecorder()

	FormPost(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	errHandler(err)

	fmt.Println(string(body))

}
