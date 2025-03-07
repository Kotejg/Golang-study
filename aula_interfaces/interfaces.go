package main

import (
	"fmt"
	"math"
)

// declaracao de structs
type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

// declaracao de interfaces
type formas interface {
	area() float64
}

// declaracao de metodos
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

// declaracao de funcoes
func escreverArea(f formas) {
	fmt.Printf("A Area da forma e %0.2f\n", f.area())
}

func main() {
	c := circulo{35}
	escreverArea(c)

	r := retangulo{2, 2}
	escreverArea(r)

}
