package goweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TemplateDataMapIFStatement(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/statement_if.gohtml"))
	t.ExecuteTemplate(w, "statement_if.gohtml", map[string]interface{}{
		"Title": "Test Template with statement if",
		"Name":  "Arif",
	})
}

func ComparateValue(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title": "Test template comparator HTML",
		"Name":  "Joko",
		"Nilai": 100,
	})
}

func IterateRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Test template Iterate with range",
		"Hobbies": []string{
			"Game", "Sport", "Code",
		},
	})
}

func NestedMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/nested_map.gohtml"))
	t.ExecuteTemplate(w, "nested_map.gohtml", map[string]interface{}{
		"Name": "Arif",
		"Address": map[string]interface{}{
			"Street": "Jalan jalan kemana",
			"City":   "Namek Planet",
		},
	})
}

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))

	t.ExecuteTemplate(w, "layout.gohtml", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Arif",
	})
}

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("Function").Parse(`{{.SayHello "Dodo"}}`))
	t.ExecuteTemplate(w, "Functionw", MyPage{
		Name: "Arif",
	})
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("Functionw").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(w, "Functionw", MyPage{
		Name: "Dading",
	})
}

func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("Functions")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{upper .Name}}`))
	t.ExecuteTemplate(w, "Functions", MyPage{
		Name: "Arif Rachman Hakim",
	})
}

//go:embed templates/*.gohtml
var templated embed.FS
var myTemplated = template.Must(template.ParseFS(templated, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplated.ExecuteTemplate(w, "simple.gohtml", "Hello Custom template caching")
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

func TestTemplateDataMapStatementIF(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	TemplateDataMapIFStatement(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestCompareValue(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	ComparateValue(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestIterateRangeValue(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	IterateRange(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestNestedMap(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	NestedMap(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestTemplateLayout(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	TemplateLayout(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestTemplateFunction(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost"+port, nil)
	rec := httptest.NewRecorder()

	TemplateFunction(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestTemplateFunctionGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	TemplateFunctionGlobal(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}

func TestTemplateCaching(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:"+port, nil)
	rec := httptest.NewRecorder()

	TemplateCaching(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	errHandler(err)
	fmt.Println(string(body))
}
