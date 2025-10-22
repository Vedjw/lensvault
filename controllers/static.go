package controllers

import (
	"html/template"
	"net/http"

	"github.com/Vedjw/lensvault/templates"
	"github.com/Vedjw/lensvault/views"
)

func StaticHandler(pattern ...string) http.HandlerFunc {
	tpl := views.Must(views.ParseFS(templates.FS, pattern...))
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(pattern ...string) http.HandlerFunc {
	tpl := views.Must(views.ParseFS(templates.FS, pattern...))
	questions := []struct {
		Q string
		A template.HTML
	}{
		{
			Q: "Is there a free version?",
			A: "Yes! We offer a free trial for 30 days on any paid plans.",
		},
		{
			Q: "What are your support hours?",
			A: `We have support staff answering emails 24/7, though response
			times may be a bit slower on weekends.`,
		},
		{
			Q: "How do I contact support?",
			A: `Email us - <a href="mailto:support@lensvault.com">support@lensvault.com</a>`,
		},
		{
			Q: "Where is your office?",
			A: "Our entire team is remote!",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
