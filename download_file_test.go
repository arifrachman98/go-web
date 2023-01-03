package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")

	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}

	// w.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"") function force download
	http.ServeFile(w, r, "./resources/"+file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	errHandler(err)
}
