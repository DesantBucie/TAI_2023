package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

func pageFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["mode"]
	id := vars["id"]
	fmt.Fprintf(w, "<html><body>PG: %v<br>ID: %v</body></html>", page, id)
}

func losujFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a, err := strconv.Atoi(vars["a"])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(vars["b"])
	if err != nil {
		panic(err)
	}
	if b < a {
		fmt.Fprintf(w, "<html><body><label style='color:red;'>B is smaller than A</label></body></html>")
	} else {
		c := rand.Intn(b-a+1) + a
		fmt.Fprintf(w, "<html><body>Liczba: %v</body></html>", c)
	}
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/losuj/{a:[0-9]+}/{b:[0-9]+}", losujFunc)
	r.HandleFunc("/page/{mode}/{id:[0-9]+}", pageFunc)
	http.ListenAndServe(":8080", r)
}
