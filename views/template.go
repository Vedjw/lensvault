package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

// ^Only use this func in main, generally
func Must(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseTpl(fs fs.FS, patterns ...string) (*Template, error) {
	tpl, err := template.ParseFS(fs, patterns...)
	if err != nil {
		return &Template{}, fmt.Errorf("parsing template through FS: %w", err)
	}
	return &Template{
		htmlTpl: tpl,
	}, nil
}

func (t *Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template",
			http.StatusInternalServerError)
		return
	}
}
