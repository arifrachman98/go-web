package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hell yeah")
	} else {
		fmt.Fprintf(w, "Hai %s", name)
	}
}

func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestQueryParameter(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port+"/hello?name=Arif", nil)
	rec := httptest.NewRecorder()

	SayHello(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestMultipleQueryParameter(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port+"/hello?first_name=arif&last_name=rachman", nil)
	rec := httptest.NewRecorder()

	MultipleQueryParameter(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	errHandler(err)
	fmt.Println(string(body))

}
