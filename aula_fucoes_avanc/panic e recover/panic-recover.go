package main

import "fmt"

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperaFunc()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("media exatamente 6!")
}

//recupera
func recuperaFunc() {
	if r := recover(); r != nil {
		fmt.Println("execucao de recupecao com sucesso!")
	}
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("pós execucao")
}
