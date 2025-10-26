package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // tambem suportado no switch
		fmt.Println("Você ganhou!!!")
		fmt.Println("Com o número:", i)
	} else {
		fmt.Println("Você perdeu!!!")
		fmt.Println("Com o número:", i)
	}
}
