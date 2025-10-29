package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 2000,
			"Gustavo":  3000,
			"Gabriel":  5000,
		},
		"J": {
			"João": 6500,
			"José": 2500,
			"Juan": 4500,
		},
		"P": {
			"Pedro":   7000,
			"Paulo":   6000,
			"Patrick": 1500,
		},
	}

	delete(funcsPorLetra, "J")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
