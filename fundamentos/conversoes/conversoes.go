package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println("A nota final é", notaFinal)

	// cuidado...
	fmt.Println("\n----- Cuidado... -----")
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("\n----- Int para String -----")
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	fmt.Println("\n----- String para Int -----")
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// convertendo boolean
	fmt.Println("\n----- Boolean -----")
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
