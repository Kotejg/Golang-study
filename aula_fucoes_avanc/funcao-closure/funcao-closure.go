package main

import "fmt"

func closure() func() {
	txt := "dentro da funcao closure"

	funcao := func() {
		fmt.Println(txt)
	}

	return funcao
}

func main() {
	txt := "dentro da funcao main"
	fmt.Println(txt)

	funcaoNova := closure()

	funcaoNova()

}
