package main

import "fmt"

func inc1(n int) {
	n++ // n = n + 1
}

// Revisão: um ponterio é representado por um *
func inc2(n *int) {
	// Revisão: * é usado para desrefenciar, ou seja
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// Revisão: & usado para obter o endereço da variável
	inc2(&n) // por referência
	fmt.Println(n)
}
