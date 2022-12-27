package goweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	t := template.Must(template.New("Simple").Parse(templateText))

	t.ExecuteTemplate(w, "Simple", "Test Template HTML on Go-Lang")
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Test File HTML Template")
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Test Directory HTML Template")
}

//go:embed templates
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Test Embed HTML Template")
}

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Template data Map",
		"Name":  "Arif",
	})
}

type Page struct {
	Title string
	Name  string
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template data struct",
		Name:  "Arif",
	})
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

func TestSimpleHTMLDirectory(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	TemplateDirectory(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestSimpleHTMLEmbed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	TemplateEmbed(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestTemplateDataMap(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	TemplateDataMap(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestTemplateDataStruct(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	TemplateDataStruct(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}
