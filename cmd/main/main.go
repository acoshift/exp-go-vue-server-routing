package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const templateDir = "templates"

var templates = map[string]*template.Template{}

func joinTemplateDir(files []string) []string {
	r := make([]string, len(files))
	for i, f := range files {
		r[i] = filepath.Join(templateDir, f+".html")
	}
	return r
}

func parseTemplates(set []string) {
	templateName := set[0]
	t := template.New("")
	t.Funcs(template.FuncMap{
		"templateName": func() string { return templateName },
		"json": func(v interface{}) string {
			r, _ := json.Marshal(v)
			return string(r)
		},
	})
	_, err := t.ParseFiles(joinTemplateDir(set)...)
	must(err)
	t = t.Lookup("root")
	if t == nil {
		log.Fatalf("root template not found in %v", set)
	}
	templates[templateName] = t
}

var templateEntrypoint = [][]string{
	{"index", "layout"},
	{"list", "layout"},
	{"view", "layout"},
}

func executeTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	if len(name) == 0 {
		name = "index"
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if t := templates[name]; t != nil {
		err := t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.NotFound(w, r)
}

func main() {
	for _, p := range templateEntrypoint {
		parseTemplates(p)
	}
	must(http.ListenAndServe(":8080", http.HandlerFunc(handler)))
}

type book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var books = []*book{
	{"1", "Book 1", "Book 1 Description"},
	{"2", "Book 2", "Book 2 Description"},
	{"3", "Book 3", "Book 3 Description"},
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := path.Clean(r.URL.Path)
	log.Println(p)
	p = strings.TrimPrefix(p, "/")
	if strings.HasPrefix(p, "-/components/") {
		http.ServeFile(w, r, "components/"+strings.TrimPrefix(p, "-/components/"))
		return
	}
	var data interface{}
	if p == "list" {
		data = map[string]interface{}{
			"books": books,
		}
	} else if p == "view" {
		id := r.URL.Query().Get("id")
		for _, x := range books {
			if x.ID == id {
				data = x
				break
			}
		}
	}

	executeTemplate(w, r, p, data)
}
