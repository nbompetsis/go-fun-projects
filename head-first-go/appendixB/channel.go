package main

import (
	"fmt"
	"time"
)

func sendLetters(channel chan string) {
	time.Sleep(time.Second)
	channel <- "a"
	time.Sleep(time.Second)
	channel <- "b"
	time.Sleep(time.Second)
	channel <- "c"
	time.Sleep(time.Second)
	channel <- "d"
}

func main() {
	fmt.Println(time.Now())
	channel := make(chan string, 3)

	go sendLetters(channel)
	time.Sleep(5 * time.Second)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(time.Now())
}
