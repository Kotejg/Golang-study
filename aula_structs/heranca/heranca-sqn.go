package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	pessoa
	matricula uint8
}

func main() {
	fmt.Println("algo proximo de heranca")
	//ao inves de estudante.pessoa.idade
	// estudante.idade
}
