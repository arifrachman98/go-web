package goweb

import (
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

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	errHandler(err)
}
