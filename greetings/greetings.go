package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages.
func randomFormat() string {
	formats := []string{
		"Hi, %s. Welcome!",
		"Great to see you, %s!",
		"Hail, %s! Well met!",
		"Hello, %s! Welcome to my App.",
	}
	return formats[rand.Intn(len(formats))]
}
