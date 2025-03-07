package main

import "fmt"

func main() {
	fmt.Println("estrutura de controle")

	num := 10

	if num > 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("menor que 15")
	}

	if outroNumero := num; outroNumero > 0 {
		fmt.Println("numero e maior que 0")
	} else if num < 10 {
		fmt.Println("menos q 10")
	} else {
		fmt.Println("maior q 10")
	}
}
