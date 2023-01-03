package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Before Middleware")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After MiddleWare")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(w, "Welcome Middleware")
	})

	mux.HandleFunc("/another", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(w, "This another page")
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		panic("Slow")
	})

	logMiddleware := LogMiddleware{
		Handler: mux,
	}

	errorHandler := ErrorHandler{
		Handler: &logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: &errorHandler,
	}

	err := server.ListenAndServe()
	errHandler(err)
}
