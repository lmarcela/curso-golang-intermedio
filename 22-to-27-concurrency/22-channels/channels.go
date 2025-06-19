package main

import "fmt"

func main() {
	// 1- Unbuffered channels
	// Unbuffered channels block the sender until the receiver is ready to read from the channel
	// and vice versa. This means that if you try to send data on an unbuffered channel without
	// a receiver ready to read from it, your program will deadlock.
	// ch := make(chan int)
	// ch <- 1
	// fmt.Println(<-ch)

	//2- Buffered channels
	// Buffered channels allow you to send data without blocking the sender until the buffer is full.
	// If the buffer is full, the sender will block until there is space in the buffer.
	// This means that you can send data on a buffered channel without a receiver ready to read
	// from the channel, as long as the buffer has space available.
	// However, if you try to send more data than the buffer can hold, your program
	// will still block until there is space in the buffer.
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	// c <- 3 // This will block the sender until there is space in the buffer

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// fmt.Println(<-c) // This will block the receiver until there is data in the channel
}
