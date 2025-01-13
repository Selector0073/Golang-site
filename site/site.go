package main

import (
	"html/template"
	"log"
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

type orderedList struct {
	Info1 string
	Info2 string
	Info3 string
}

func home_page(w http.ResponseWriter, r *http.Request) {
	lang1 := language{"Golang", "Python", "Compiler", "Interpreter", "Schnell", "Langsam"}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, lang1)
}

func more_page(w http.ResponseWriter, r *http.Request) {
	List := orderedList{"Schnell", "Komfortabel", "Fortschrittlich"}

	tmpl, _ := template.ParseFiles("templates/sec.html")
	tmpl.Execute(w, List)

}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/more/", more_page)
	log.Println("Der Server läuft auf Port 8080")
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
