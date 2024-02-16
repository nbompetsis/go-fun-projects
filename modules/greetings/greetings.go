package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func HelloFunction(name string) (string, error) {
	if name == "" {
		return "", errors.New("The name should not be null")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, value := range names {
		message, err := HelloFunction(value)
		if err != nil {
			return nil, err
		}
		messages[value] = message
	}

	return messages, nil
}

func randomFormat() string {
	var randomFormats = []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return randomFormats[rand.Intn(len(randomFormats))]
}
