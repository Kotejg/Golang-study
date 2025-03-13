package main

import (
	"fmt"
	"time"
)

func main() {
	//concorrente != paralelismo
	go escrever("Ola Mundo")
	escrever("Jefff")
}

func escrever(txt string) {
	for {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}
