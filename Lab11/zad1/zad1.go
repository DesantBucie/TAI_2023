package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Student struct {
	FirstName string
	LastName  string
	Index     int
	Mail      string
}

var studenci []Student = []Student{
	{"Jan", "Kowalski", 12345, "test@test"},
	{"Marek", "Nowak", 30000, "to@tamto"},
	{"Anna", "Zdyb", 23232, "anna@zdyb"},
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "<html><body>STRONA GŁOWNA</body></html>")
}

func daneFunc(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("../pages/dane.html")
	tmpl.Execute(w, studenci)
}
func main() {
	http.HandleFunc("/test/", daneFunc)
	http.ListenAndServe("localhost:8080", nil)
}
