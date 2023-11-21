package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func chargeValue(x int, y int) error {
	if x > y {
		return errors.New("x>y")
	} else {
		return errors.New("x<=y")
	}
}

func createPanic() {
	panic(errors.New("this is panic"))
}

func callCreatePanic() {
	defer func() {
		fmt.Println("call create panic")
	}()
	createPanic()
}

func hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func main() {
	err := chargeValue(1, 2)
	fmt.Println(err)

	switch err := err.(type) {
	case *os.PathError:
		fmt.Println("os path error")
	default:
		fmt.Println(err)
	}

	defer func() {
		fmt.Println("this is main")
	}()

	callCreatePanic()

	message, err := hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
