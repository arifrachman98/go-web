package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	// t, err := template.New("SIMPLE").Parse(templateText)
	// errHandler(err)

	t := template.Must(template.New("Simple").Parse(templateText))

	t.ExecuteTemplate(w, "Simple", "Test Template HTML on Go-Lang")
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Test Template")
}

func TestSimpleHTML(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	SimpleHTML(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestSimpleHTMLFile(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	SimpleHTMLFile(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}
