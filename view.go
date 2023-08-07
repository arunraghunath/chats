package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Template struct {
	viewts *template.Template
}

func ParseTemplate(files ...string) *Template {
	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Print(err)
		panic(err)

	}
	return &Template{
		viewts: ts,
	}
}

func (ts *Template) Execute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	err := ts.viewts.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error Parsing the web page", http.StatusInternalServerError)
		return
	}
}
