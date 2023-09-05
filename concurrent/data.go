package main

import "fmt"

func sender(value int, channel chan int) {
	channel <- value
}

func main() {
	channel := make(chan int)
	go sender(1, channel)
	go sender(2, channel)
	x, y := <-channel, <-channel

	fmt.Printf("X = %d Y = %d \n", x, y)
}
