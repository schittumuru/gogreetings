package gogreetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"rsc.io/quote"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func HelloRandom() string {
	return quote.Go()
}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, name := range names {

		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	fmt.Println("Initializing the greetings module!")
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {

	formats := []string{
		"Hi %v, welcome!",
		"Yo %v, wassup",
		"Greeting Mr %v.",
	}

	return formats[rand.Intn(len(formats))]
}
