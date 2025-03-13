package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Ola mundo!", canal)
	fmt.Println("Depois da funcoa escrever ser executada")
	//for {
	//	mensagem, verificaCanal := <-canal
	//	if !verificaCanal {
	//		fmt.Println(verificaCanal)
	//		break
	//	}
	//	fmt.Println(mensagem)
	//}
	//enquanto o canal estiver aberto, ele executa o loop
	for message := range canal {
		fmt.Println(message)
	}

	fmt.Println("fim do programa")

}

func escrever(txt string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- txt
		time.Sleep(time.Second)
	}

	close(canal)
}
