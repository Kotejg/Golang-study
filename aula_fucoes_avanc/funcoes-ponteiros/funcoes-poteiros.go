package main

import (
	"fmt"
)

func inverteSinal(numero int) int {
	return numero * -1
}

func inverteSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	fmt.Println(numero)
	fmt.Println(inverteSinal(numero))

	novoNum := 40
	fmt.Println("numero antigo")
	fmt.Println(novoNum)
	inverteSinalComPonteiro(&novoNum)
	fmt.Println("valor alterado por funcao")
	fmt.Println(novoNum)
}
