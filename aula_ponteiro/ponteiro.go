package main

import "fmt"

func main() {

	var vari1 int = 10
	var vari2 int = vari1

	fmt.Println(vari1, vari2)

	vari1++

	fmt.Println(vari1, vari2)

	//ponteiro e uma refencia de memoria

	var vari3 int
	var varPonteiro *int

	vari3 = 100
	varPonteiro = &vari3

	fmt.Println(vari3, varPonteiro) // desreferenciacao
}
