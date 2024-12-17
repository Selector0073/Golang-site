package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type language struct {
	Name  string
	Name2 string
	Tpe   string
	Tpe2  string
	More  string
	More2 string
}

func home_page(w http.ResponseWriter, r *http.Request) {
	lang1 := language{"Golang", "Python", "Compiler", "Interpreter", "Schnell", "Langsam"}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, lang1)
}

func about_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yes Rico kaboom!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/about/", about_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
