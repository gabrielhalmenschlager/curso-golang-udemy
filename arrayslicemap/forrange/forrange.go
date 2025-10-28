package main

import "fmt"

func main() {
	// numeros := [] int{1, 2, 3, 4, 5} sem os tres pontos seria um slice
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Printf("%d\n", num)
	}
}
