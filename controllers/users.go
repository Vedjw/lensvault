package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Vedjw/lensvault/models"
)

type Users struct {
	Templates struct {
		NewUser Template
		Signin  Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// var data struct {
	// 	Email string
	// }
	// data.Email = r.FormValue("email")
	u.Templates.NewUser.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	ageStr := strings.TrimSpace(r.FormValue("age"))
	age, err := strconv.Atoi(ageStr)
	if err != nil || age < 12 || age > 100 {
		http.Error(w, "Invalid age value", http.StatusBadRequest)
		return
	}
	nu := &models.NewUser{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Age:       age,
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}

	user, err := u.UserService.Create(nu)
	if err != nil {
		fmt.Printf("controllers: create user : %v\n", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created %+v", user)
}

func (u Users) Signin(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.Signin.Execute(w, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User authenticated: %+v", user)
}
