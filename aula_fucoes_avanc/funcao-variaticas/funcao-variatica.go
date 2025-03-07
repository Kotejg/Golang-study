package main

import "fmt"

func somar(numeros ...int) int {
	total := 0

	for _, num := range numeros {
		total += num
	}

	return total
}

func escrever(txt string, numeros ...int) {
	for _, num := range numeros {
		fmt.Println(txt, num)
	}

}

func main() {

	fmt.Println(somar(1, 2, 3, 4, 5))

	escrever("ola ", 1, 2, 3, 4, 5)
}
