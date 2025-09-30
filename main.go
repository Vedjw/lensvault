package main

import (
	"fmt"
	"net/http"

	"github.com/Vedjw/lensvault/controllers"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler("home.gohtml", "tailwind.gohtml"))

	r.Get("/contact", controllers.StaticHandler("contact.gohtml", "tailwind.gohtml"))

	r.Get("/faq", controllers.FAQ("faq.gohtml", "tailwind.gohtml"))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting Server on :3000")
	http.ListenAndServe(":3000", r)
}
