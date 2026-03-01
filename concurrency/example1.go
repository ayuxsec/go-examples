package main

import (
	"fmt"
	"time"
)

func main() {
	num1 := 5
	num2 := 6
	go add(num1, num2)
	time.Sleep(3 * time.Second) // wait for goroutine to finish
}

func add(x, y int) {
	fmt.Println(x + y) // 11
}
