package main

import "fmt"

func odd(channel chan int) {
	channel <- 1
	channel <- 3
}

func even(channel chan int) {
	channel <- 2
	channel <- 4
}

func main() {
	channelOdd := make(chan int)
	channelEven := make(chan int)

	go odd(channelOdd)
	go even(channelEven)

	fmt.Println(<-channelOdd)
	fmt.Println(<-channelEven)
	fmt.Println(<-channelOdd)
	fmt.Println(<-channelEven)
}
