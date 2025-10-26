package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarIphone15 := trab1 && trab2
	comprarIphone13 := trab1 != trab2 // ou exclusivo
	comprarAcai := trab1 || trab2

	return comprarIphone15, comprarIphone13, comprarAcai
}

func main() {
	iph15, iph13, acai := compras(true, true)
	fmt.Printf("Iphone 15: %t\nIphone 13: %t\nAçai: %t\nSaudável: %t",
		iph15, iph13, acai, !acai)
}
