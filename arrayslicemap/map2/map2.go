package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José":     10000,
		"Maria":    5000,
		"Gabriela": 2000,
	}

	funcsESalarios["Cristiane"] = 3500
	delete(funcsESalarios, "Inexistente") // se tentar excluir um elemento que não existe não tem problema nenhum

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
