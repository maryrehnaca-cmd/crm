package handlers

import (
	"crm/db"
	"html/template"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/login.html"))
	t.Execute(w, nil)
}

func SignupPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/signup.html"))
	t.Execute(w, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	name := r.FormValue("fullname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err := db.DB.Exec(
		"INSERT INTO users(name, email, password) VALUES($1,$2,$3)",
		name, email, password,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	var userID int

	err := db.DB.QueryRow(
		"SELECT user_id FROM users WHERE email=$1 AND password=$2",
		email,
		password,
	).Scan(&userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
