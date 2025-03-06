package main

import (
	"aula_pacotes/auxiliar"
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("escrevendo no arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("eeeecom")
	fmt.Println(erro)
}
