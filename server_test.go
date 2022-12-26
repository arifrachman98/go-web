package goweb

import (
	"net/http"
	"testing"
)

func errHandler(err error)  {
	if err != nil {
		panic(err)
	}
}

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	errHandler(err)
}