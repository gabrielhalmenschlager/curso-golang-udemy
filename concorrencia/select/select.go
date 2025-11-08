package main

import (
	"fmt"
	"time"

	"github.com/gabrielhalmenschlager/curso-golang-udemy/html"
)

func oMaisRapido(ur1, ur2, ur3 string) string {
	c1 := html.Titulo(ur1)
	c2 := html.Titulo(ur2)
	c3 := html.Titulo(ur3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://chatgpt.com/",
		"https://www.youtube.com",
		"https://www.google.com",
	)
	fmt.Println(campeao)
}
