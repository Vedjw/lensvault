package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		NewUser Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.NewUser.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Email: ", r.FormValue("email"))
	fmt.Fprintln(w)
	fmt.Fprint(w, "Password: ", r.FormValue("password"))
}
