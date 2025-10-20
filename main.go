package main

import (
	"fmt"
	"net/http"

	"github.com/Vedjw/lensvault/controllers"
	"github.com/Vedjw/lensvault/models"
	"github.com/Vedjw/lensvault/templates"
	"github.com/Vedjw/lensvault/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler("home.gohtml", "tailwind.gohtml"))

	r.Get("/contact", controllers.StaticHandler("contact.gohtml", "tailwind.gohtml"))

	r.Get("/faq", controllers.FAQ("faq.gohtml", "tailwind.gohtml"))

	config := models.DefaultPostgresConfig()
	db, err := models.Open(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	userService := models.UserService{
		DB: db,
	}
	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.NewUser = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	usersC.Templates.Signin = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", "tailwind.gohtml",
	))
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)
	r.Get("/signin", usersC.Signin)
	r.Post("/signin", usersC.ProcessSignIn)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting Server on :3000")
	http.ListenAndServe(":3000", r)
}
