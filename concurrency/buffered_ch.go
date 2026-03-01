package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1) // buffer size 1

	go func() {
		time.Sleep(2 * time.Second)
		c <- 10 // does not block (buffer has space)
	}()

	fmt.Println("doing other work...")
	time.Sleep(1 * time.Second)

	result := <-c // waits if value not ready yet
	fmt.Println(result)
}
