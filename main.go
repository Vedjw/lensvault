package main

import (
	"fmt"
	"net/http"

	"github.com/Vedjw/lensvault/controllers"
	"github.com/Vedjw/lensvault/templates"
	"github.com/Vedjw/lensvault/views"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseTpl(templates.FS, "home.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseTpl(templates.FS, "contact.gohtml"))))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.ParseTpl(templates.FS, "faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting Server on :3000")
	http.ListenAndServe(":3000", r)
}
