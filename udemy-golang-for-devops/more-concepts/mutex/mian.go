package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MyType struct {
	Counter int
	mu      sync.Mutex
}

func main() {
	myTypeInstance := MyType{}
	finished := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(myTypeInstance *MyType, c chan bool) {
			// mutex exclusion
			myTypeInstance.mu.Lock()
			myTypeInstance.Counter += 1
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			if myTypeInstance.Counter == 5 {
				fmt.Printf("Found counter == 5  \n")
			}
			myTypeInstance.mu.Unlock()
			// mutex exclusion
			c <- true
		}(&myTypeInstance, finished)
	}

	for i := 0; i < 5; i++ {
		<-finished
	}

	fmt.Printf("Counter: %d\n", myTypeInstance.Counter)

}
