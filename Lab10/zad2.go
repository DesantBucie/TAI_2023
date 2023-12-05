package main

import "fmt"

func minmax(tab []int) (int, int) {
	if len(tab) < 1 {
		return 0, 0
	}
	min, max := tab[0], tab[0]
	for i := 0; i < len(tab); i++ {
		if min > tab[i] {
			min = tab[i]
		}
		if max < tab[i] {
			max = tab[i]
		}
	}
	return min, max
}
func main() {
	numbers := []int{10, 2, 24, 13, 20}
	a, b := minmax(numbers)
	fmt.Println("Min: ", a, "Max: ", b)
}
