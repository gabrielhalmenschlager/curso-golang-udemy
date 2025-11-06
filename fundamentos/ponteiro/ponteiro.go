package main

import "fmt"

/*
Um ponteiro em Go é uma variável que armazena o endereço de memória
de outra variável, e não o valor em si.
*/

func main() {
	i := 1

	var p *int = nil
	p = &i // pegando o endereço da variável
	*p++   // desreferenciando (pegando o valor)
	i++

	// Go não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
