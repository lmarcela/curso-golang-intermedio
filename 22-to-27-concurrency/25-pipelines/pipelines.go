package main

import "fmt"

func Generator(c chan<- int) { // chan<- = Write-only channel
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) { // <-chan = Read-only channel, chan<- = Write-only channel
	for valueIn := range in {
		out <- 2 * valueIn
		//in <- 1
	}
	close(out)
}

func Print(c <-chan int) { // <-chan = Read-only channel
	for valueC := range c {
		fmt.Println(valueC)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
}
