package main

import "fmt"

func main() {
	// homogênea (mesmo tipo) e estática (fixo)

	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 8.5, 4.2, 6.3
	//notas[3] = 10
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Media %.2f\n", media)
}
