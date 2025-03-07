package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma, subs int) {
	soma = n1 + n2
	subs = n1 - n2
	return
}

func main() {
	soma, substracao := calculosMatematicos(10, 15)

	fmt.Println(soma, substracao)
}
