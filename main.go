package main

import (
	"fmt"
	"net/http"
)

func data(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func main() {
	http.HandleFunc("/", data)
	fileHandler := http.FileServer(http.Dir("templates/static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileHandler))
	fmt.Println("Server starting at :2020")
	http.ListenAndServe(":2020", nil)
}
