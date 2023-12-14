package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

type jsonIPdata struct { // stuktura danych do której przekażemy dane JSON
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countrycode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionname"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
}

func daneFunc(w http.ResponseWriter, r *http.Request, ipData jsonIPdata) {
	tmpl, _ := template.ParseFiles("../pages/zad5.html")
	tmpl.Execute(w, ipData)
}

func main() {
	client := http.Client{Timeout: time.Second}       // utworzenie klienta HTTP
	url := "http://ip-api.com/json/87.99.119.110"     // strona z API z danych
	r, e := http.NewRequest(http.MethodGet, url, nil) // żądanie HTTP GET
	if e != nil {
		log.Fatal(e)
	}
	res, err := client.Do(r) // wywołanie żądania HTTP GET
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body) // odczytanie odpowiedzi z HTTP GET
	if err != nil {
		log.Fatal(err)
	}
	ipData := jsonIPdata{}              // utworzenie struktury na otrzymane dane
	err = json.Unmarshal(body, &ipData) // parsowanie danych JSON na strukturę
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/87.99.119.110/", func(w http.ResponseWriter, r *http.Request) {
		daneFunc(w, r, ipData)
	})
	http.ListenAndServe(":8080", nil)
	fmt.Println(ipData) // wyświetlenie wyniku w terminalu
}
