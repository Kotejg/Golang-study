package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica(12)
	generica("dsdsds")
	generica(true)

	fmt.Println()
}
