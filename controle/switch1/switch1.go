package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println("Sua nota é", notaParaConceito(9))
	fmt.Println("Sua nota é", notaParaConceito(8))
	fmt.Println("Sua nota é", notaParaConceito(5))
	fmt.Println("Sua nota é", notaParaConceito(3))
	fmt.Println("Sua nota é", notaParaConceito(1))
	fmt.Println("Sua nota é", notaParaConceito(11))
}
