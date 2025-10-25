package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Vedjw/lensvault/controllers"
	"github.com/Vedjw/lensvault/models"
	"github.com/Vedjw/lensvault/templates"
	"github.com/Vedjw/lensvault/views"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func logMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		method := r.Method
		switch method {
		case "GET":
			fmt.Printf("%s      %s   %v\n", r.Method, r.RequestURI, elapsed)
		case "POST":
			fmt.Printf("%s     %s   %v\n", r.Method, r.RequestURI, elapsed)
		case "DELETE":
			fmt.Printf("%s  %s   %v\n", r.Method, r.RequestURI, elapsed)
		case "PUT":
			fmt.Printf("%s      %s   %v\n", r.Method, r.RequestURI, elapsed)
		}
	})
}

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
	sessionService := models.SessionService{
		DB:            db,
		BytesPerToken: 32,
	}
	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
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
	r.Post("/signout", usersC.ProcessSignOut)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		//TODO: fix this befor deployment
		csrf.Secure(false),
		csrf.TrustedOrigins([]string{"localhost:3000"}),
	)

	fmt.Println("Starting Server on :3000")
	http.ListenAndServe(":3000", logMw(csrfMw(r)))
}
