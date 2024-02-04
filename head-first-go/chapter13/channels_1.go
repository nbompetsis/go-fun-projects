package main

import (
	"fmt"
	"time"
)

func greetings(channel chan string) {
	fmt.Println("Waiting a second")
	time.Sleep(time.Second)
	channel <- "Geia sou"
}

func main() {
	channel := make(chan string)

	go greetings(channel)
	fmt.Println(<-channel)

	// channel <- "Geia sou deadlock" // Deadlock
}
