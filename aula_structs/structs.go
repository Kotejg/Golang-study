package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco
}

type endereco struct {
	numero uint8
}

func main() {
	fmt.Println("arquivo struct")

	enderecoExemplo := endereco{5}

	//1 forma
	var user usuario
	user.idade = 18
	user.nome = "Nanda"
	fmt.Println(user)
	//2 forma
	user2 := usuario{"davi", 1, enderecoExemplo}

	fmt.Println(user2)

	//3 forma sem passar todos os atributos
	user3 := usuario{nome: "Jeff"}
	fmt.Println(user3)
}
