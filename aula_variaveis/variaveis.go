package main

import "fmt"

func main() {
	var variavelTxt string = "Jeff"
	variavelTxt2 := "Ana"

	fmt.Println(variavelTxt2)
	fmt.Println(variavelTxt)

	var (
		abc string = "Joao"
		cde string = "Math"
	)
	fmt.Println(abc, cde)

	vari4, vari5 := "nanda", "vitoria"

	fmt.Println(vari4, vari5)

	vari4, vari5 = vari5, vari4

	fmt.Println(vari4, vari5)

}
