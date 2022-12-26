package goweb

import (
	"net/http"
	"testing"
)

var port = "8080"

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:" + port,
	}

	err := server.ListenAndServe()
	errHandler(err)
}
