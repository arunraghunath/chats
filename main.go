package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func data(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func signup(w http.ResponseWriter, r *http.Request) {
	signupTmpl, err := template.ParseFiles("templates/signup.html")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	err = signupTmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func submit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse the form", http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	fullname := r.FormValue("name")
	password := r.FormValue("pwd")
	fmt.Fprintf(w, "Email is %s, Fullname is %s, Password is %s", email, fullname, password)
}

func main() {
	http.HandleFunc("/", data)
	fileHandler := http.FileServer(http.Dir("templates/static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileHandler))
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/submit", submit)
	fmt.Println("Server starting at :2020")
	http.ListenAndServe(":2020", nil)
}
