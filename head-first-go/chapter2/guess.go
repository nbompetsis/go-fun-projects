// guess challenges players to guess a random number.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const MIN, MAX = 1, 100
const MAX_GUESSES = 10

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	target := random.Intn(MAX) + MIN
	success := false
	reader := bufio.NewReader(os.Stdin)

	for guesses := 0; guesses < MAX_GUESSES; guesses++ {
		fmt.Printf("Choose a number between [%d, %d], you have %d guesses left\n", MIN, MAX, (MAX_GUESSES - guesses))
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Could not parse user's input")
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal("Could not parse user's input")
		}

		if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		} else if guess < target {
			fmt.Println("Oops. Your guess was LOW")
		} else {
			success = true
			fmt.Println("Perfect, you guessed it!")
			break
		}
	}

	if !success {
		fmt.Printf("\nGame Over\nSorry. You didn’t guess my number. It was: [%d]\n", target)
	}
}
