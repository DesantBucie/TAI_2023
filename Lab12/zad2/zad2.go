package main

import (
	"fmt"
	"github.com/gorilla/mux"
	deep "github.com/patrikeh/go-deep"
	"github.com/patrikeh/go-deep/training"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var n *deep.Neural

var data = training.Examples{
	{Input: []float64{2.7810836, 2.550537003}, Response: []float64{1}},
	{Input: []float64{1.465489372, 2.362125076}, Response: []float64{0}},
	{Input: []float64{3.396561688, 4.400293529}, Response: []float64{0}},
	{Input: []float64{1.38807019, 1.850220317}, Response: []float64{0}},
	{Input: []float64{7.627531214, 2.759262235}, Response: []float64{1}},
	{Input: []float64{5.332441248, 2.088626775}, Response: []float64{1}},
	{Input: []float64{6.922596716, 1.77106367}, Response: []float64{1}},
	{Input: []float64{8.675418651, -0.242068655}, Response: []float64{1}},
}

func main() {
	// inicjalizacja ziarna losowania
	rand.Seed(time.Now().UTC().UnixNano())
	// dane uczące
	// utworzenie sieci neuronowej (2 wejścia, 2-2-1 neuronów)
	n = deep.NewNeural(&deep.Config{
		Inputs:     2,
		Layout:     []int{2, 2, 1},
		Activation: deep.ActivationSigmoid,
		Mode:       deep.ModeBinary,
		Weight:     deep.NewNormal(1.0, 0.0),
		Bias:       true,
	})
	// optymalizacja sieci neuronowej
	optimizer := training.NewSGD(0.05, 0.1, 1e-6, true)
	trainer := training.NewTrainer(optimizer, 50)
	training, heldout := data.Split(0.5)
	trainer.Train(n, training, heldout, 1000)
	// wyświetlenie predykcji dla danych
	blad := 0.0
	for i := 0; i < len(data); i++ {
		predykcja := n.Predict(data[i].Input)
		prawidlowo := data[i].Response
		blad += math.Abs(predykcja[0] - prawidlowo[0])
		fmt.Println("PREDYKCJA:", predykcja, ", PRAWIDŁOWO:", prawidlowo)
		fmt.Println(", BLAD:", blad)
	}
	fmt.Println("BŁĄD (SUMA):", blad)

	r := mux.NewRouter()
	r.HandleFunc("/predykcja/{input1:[0-9]+}/{input2:[0-9]+}", handlerFunc)
	r.HandleFunc("/resetuj", resetuj)
	http.ListenAndServe(":8080", r)
}

func resetuj(w http.ResponseWriter, r *http.Request) {
	n = deep.NewNeural(&deep.Config{
		Inputs:     2,
		Layout:     []int{2, 2, 1},
		Activation: deep.ActivationSigmoid,
		Mode:       deep.ModeBinary,
		Weight:     deep.NewNormal(1.0, 0.0),
		Bias:       true,
	})
	optimizer := training.NewSGD(0.05, 0.1, 1e-6, true)
	trainer := training.NewTrainer(optimizer, 50)
	training, heldout := data.Split(0.5)
	trainer.Train(n, training, heldout, 1000)
	// wyświetlenie predykcji dla danych
	blad := 0.0
	for i := 0; i < len(data); i++ {
		predykcja := n.Predict(data[i].Input)
		prawidlowo := data[i].Response
		blad += math.Abs(predykcja[0] - prawidlowo[0])
		fmt.Println("PREDYKCJA:", predykcja, ", PRAWIDŁOWO:", prawidlowo)
		fmt.Println(", BLAD:", blad)
	}
	fmt.Fprintf(w, "<html><body>")
	fmt.Fprintf(w, "BŁĄD (SUMA): %v", blad)
	fmt.Fprintf(w, "</body></html>")
}
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a, _ := strconv.ParseFloat(vars["input1"], 64)
	b, _ := strconv.ParseFloat(vars["input2"], 64)
	var c float64
	if a < b {
		c = 0
	} else {
		c = 1
	}

	data := []float64{a, b}
	resp := [1]float64{c}

	pred := n.Predict(data)
	praw := resp
	blad := praw[0] - pred[0]
	fmt.Fprintf(w, "<html><body>")
	fmt.Fprintf(w, "PREDYKCJA: %v, PRAWIDŁOWO: %v", pred, praw)
	fmt.Fprintf(w, "BLAD: %v", blad)
	fmt.Fprintf(w, "</body></html>")
}
