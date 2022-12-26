package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	}

	server := http.Server{
		Addr:    "localhost:"+port,
		Handler: handler,
	}

	err := server.ListenAndServe()
	errHandler(err)
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hellow World!!")
	})

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Testing Page")
	})

	server := http.Server{
		Addr:    "localhost:"+port,
		Handler: mux,
	}

	err := server.ListenAndServe()
	errHandler(err)
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr: "localhost:"+port,
		Handler: handler,
	}

	err := server.ListenAndServe()
	errHandler(err)
}