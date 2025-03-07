package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("incrementado i ")
	// 	time.Sleep(time.Second)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("increment i", i)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Jeff", "ana", "vitoria"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	//for n roda em structs
}
