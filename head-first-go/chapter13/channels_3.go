package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(channel chan string) {

	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***    1")
	channel <- "a"
	fmt.Println("***sending value***    2")
	channel <- "b"
}

func main() {
	channel := make(chan string)
	go send(channel)
	reportNap("receiving goroutine", 5)
	fmt.Println("Receiving:", <-channel)
	fmt.Println("Receiving:", <-channel)

}
