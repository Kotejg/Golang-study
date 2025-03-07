package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// map chave-valor
	usuario := map[string]string{
		"nome":      "jeff",
		"sobrenome": "bispo",
	}
	//acessa a chave
	fmt.Println(usuario["nome"])

}
