package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Ola Mundo")
		waitGroup.Done()
	}()

	go func() {
		escrever("Jefff")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(txt string) {
	for i := 0; i < 6; i++ {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}
