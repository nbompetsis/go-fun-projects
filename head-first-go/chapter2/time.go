package main

import (
	"fmt"
	"time"
)

func main() {
	// now := time.Now()
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
}
