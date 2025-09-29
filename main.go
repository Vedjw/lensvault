package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Vedjw/lensvault/controllers"
	"github.com/Vedjw/lensvault/views"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	tpl, err := views.ParseTpl(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.ParseTpl(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.ParseTpl(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting Server on :3000")
	http.ListenAndServe(":3000", r)
}
