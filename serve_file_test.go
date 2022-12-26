package goweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbeded(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourceOk)
	} else {
		fmt.Fprint(w, resourceNotFound)
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	errHandler(err)
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: http.HandlerFunc(ServeFileEmbeded),
	}

	err := server.ListenAndServe()
	errHandler(err)
}
