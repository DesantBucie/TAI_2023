package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

func GenerateRandom() string {
	return "strona" + strconv.Itoa(rand.Intn(4))
}
func daneFunc(w http.ResponseWriter, r *http.Request) {
	strona := GenerateRandom()
	tmpl, _ := template.ParseFiles("../pages/" + strona + ".html")
	tmpl.Execute(w, nil)
}
func main() {
	http.HandleFunc("/test/", daneFunc)
	http.ListenAndServe("localhost:8080", nil)
}
