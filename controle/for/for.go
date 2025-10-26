package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	fmt.Println("\n--- Contador de 1 a 10 ---")
	for i <= 10 { // Não tem while em Go
		fmt.Printf("%d ", i)
		i++
	}

	fmt.Println("\n\n--- Contador de 0 a 20 pulando de dois em dois ---")
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\n\nMisturando...")
	fmt.Println("--- Par ou Impar de 1 a 10 ---")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("O numero", i, "é Par ")
		} else {
			fmt.Println("O numero", i, "é Impar ")
		}
	}

	fmt.Println("\n--- Laço Infinito ---")
	for {
		// laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
