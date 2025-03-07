package main

import "fmt"

func main() {

	func() {
		fmt.Println("ola mundo")
	}()

	func(txt string) {
		fmt.Println(txt)
	}("passando paramen") // declara e ja passa o paramentro

	retorno := func(txt string) string {
		return fmt.Sprintf("Recebido -> %s %d", txt, 34)
	}("opa")

	fmt.Println(retorno)

}
