package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[123456789] = "Gabriel"
	aprovados[987654321] = "Pedro"
	aprovados[246810120] = "Leo"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[246810120])
	delete(aprovados, 246810120)
	fmt.Println(aprovados[246810120])
}
