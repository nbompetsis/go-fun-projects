package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func getResource(t time.Time, finished chan bool) {
	fmt.Println(t)
	resp, err := http.Get("http://localhost:8080/ratelimit")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status Code", resp.StatusCode)
	if resp.StatusCode == http.StatusTooManyRequests {
		finished <- true
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Body:", string(body))
	finished <- false
}

func doEvery(d time.Duration, f func(time.Time, chan bool)) {
	finished := make(chan bool)
	for x := range time.Tick(d) {
		go f(x, finished)
		isFinished := <-finished
		if isFinished {
			break
		}
	}
}

func main() {
	// doEvery(50*time.Millisecond, getResource)
	finished := make(chan bool)
	for {
		go getResource(time.Now(), finished)
		isFinished := <-finished
		if isFinished {
			break
		}
	}
}
