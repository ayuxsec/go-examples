package main

import (
	"time"
	"fmt"
)

func main() {
	num1 := 5
	num2 := 6

	c := make(chan int)

	go add(num1, num2, c)

	result := <-c  // recieve from c (block until c is recieved)
	fmt.Println(result) // 11
}

func add(x, y int, c chan int) {
	time.Sleep(3 * time.Second)
	c <- x + y // send to channel
}
