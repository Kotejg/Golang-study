package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("arrays e slices")
	//arrays
	var array1 [5]int
	array1[0] = 2
	fmt.Println(array1)

	arrayStr := [4]string{"jeff", "ana", "vi", "jorge"}
	fmt.Println(arrayStr)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//slices
	//nao precisa limitar o tamanho
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice)

	fmt.Println("verifica os tipos")
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1))
	//adicionar
	slice = append(slice, 34)

	fmt.Println(slice)

	// na sintaxe de [n...m] n= inclusivo e m= exclusivo
	// ponteiro para o arrayStr
	slice2 := arrayStr[1:4]
	fmt.Println(slice2)

	// arrays internos
	//definine (tipo,tamanho,cap.max)
	fmt.Println("------------")
	slice3 := make([]float32, 10, 15)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

}
