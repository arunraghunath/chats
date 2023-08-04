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

func main() {
	http.HandleFunc("/", data)
	fileHandler := http.FileServer(http.Dir("templates/static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileHandler))
	http.HandleFunc("/signup", signup)
	fmt.Println("Server starting at :2020")
	http.ListenAndServe(":2020", nil)
}
