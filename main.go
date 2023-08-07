package main

import (
	"fmt"
	"net/http"

	"github.com/arunraghunath/chats/models"
	_ "github.com/lib/pq"
)

func data(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
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
	models.ConnectDB()
	http.HandleFunc("/", data)
	fileHandler := http.FileServer(http.Dir("templates/static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileHandler))
	viewTS := ParseTemplate("templates/signup.html")
	http.HandleFunc("/signup", viewTS.Execute)
	http.HandleFunc("/submit", submit)
	fmt.Println("Server starting at :2020")
	http.ListenAndServe(":2025", nil)
}
