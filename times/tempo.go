package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	temof := "2025-03-07 00:00:00"

	fmt.Println(temof)
}
