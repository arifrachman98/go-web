package goweb

import (
	"embed"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}

	err := server.ListenAndServe()
	errHandler(err)
}

//go:embed resources
var resource embed.FS

func TestFileServerGolangEmbeded(t *testing.T) {
	fileServer := http.FileServer(http.FS(resource))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}

	err := server.ListenAndServe()
	errHandler(err)

}
