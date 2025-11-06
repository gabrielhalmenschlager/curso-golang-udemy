package area

import "math"

// Pi é uma propoção numérica definida pela relação entre
// o perimetro de uma circuferência e seu diâmetro

const Pi = 3.1416

// Circ é responsável por calcular a área da circuferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é reponável por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
