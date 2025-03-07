package main

import "fmt"

//estruta de uma funcao
// func nome-func (params...) retorno {}
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	somar := n1 + n2
	subs := n1 - n2
	return somar, subs
}

func main() {

	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultadof := f("nanana")
	fmt.Println(resultadof)

	// retornoSomar, retornoSubs := calculosMatematicos(10, 15)

	// fmt.Println(retornoSomar, retornoSubs)

	// "_" ignora o resultdo de uma funcao
	retornoSomar, _ := calculosMatematicos(10, 15)

	fmt.Println(retornoSomar)

}
