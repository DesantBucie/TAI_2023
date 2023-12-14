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
	tmpl, _ := template.ParseFiles("../pages/zad4.html")
	tmpl.Execute(w, nil)
}
func main() {
	http.HandleFunc("/test/", daneFunc)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../assets/"))))
	http.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir("../styles"))))
	http.ListenAndServe("localhost:8080", nil)
}
