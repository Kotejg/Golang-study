package main

import (
	"fmt"
)

func func1() {
	fmt.Println("chamando funcao 1")
}

func func2() {
	fmt.Println("chamando funcao 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	fmt.Println("media calculada. resultado sera retornado")
	fmt.Println("entrando na funcao para verificar se aluno foi aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	// defer func1()
	// defer = adiar
	// func2()

	fmt.Println(alunoEstaAprovado(7, 8))

}
