package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Student struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Index     int    `json:"index"`
	Mail      string `json:"-"`
}

var studenci []Student = []Student{
	{"Jan", "Kowalski", 12345, "test@test"},
	{"Marek", "Nowak", 30000, "to@tamto"},
	{"Anna", "Zdyb", 23232, "anna@zdyb"},
}

func jsonFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(studenci)
	w.Write(data)
}

func daneFunc(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("../pages/dane.html")
	tmpl.Execute(w, studenci)
}
func main() {
	http.HandleFunc("/test/", daneFunc)
	http.HandleFunc("/json/", jsonFunc)
	http.ListenAndServe("localhost:8080", nil)
}
