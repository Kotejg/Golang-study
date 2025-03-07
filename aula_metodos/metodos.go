package main

import "fmt"

type user struct {
	nome  string
	idade uint8
}

func (u user) salvar() {
	fmt.Printf("salvando os dados do usuario %s no banco de dados\n", u.nome)
}

func (u user) verificaMaiorDeIdade() bool {
	return u.idade >= 18
}

func main() {
	user1 := user{"jeff", 25}
	fmt.Println(user1)
	user1.salvar()
	fmt.Println(user1.verificaMaiorDeIdade())
}
