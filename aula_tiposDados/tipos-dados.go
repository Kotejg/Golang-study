package main

import (
	"errors"
	"fmt"
)

func main() {
	//tipos numericos
	// tipo int pega o valor base da sua maquina
	// var (
	// 	int, int8, int16, int32, int64
	// )

	var numero int16 = 555

	fmt.Println(numero)

	//alias
	/*
		int32 = rune
		int8 = byte
	*/

	var float float32 = 32.5

	var float2 float64 = 64.555

	fmt.Println(float, float2)

	var texto string
	fmt.Println(texto)

	txt := "jeff"
	fmt.Println(txt)

	var boolean bool = true
	var booleanf bool
	fmt.Println(booleanf, boolean)

	var erro error = errors.New("Erro novo")
	fmt.Println(erro)
}
