package main

import (
	"errors"
	"fmt"
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

}
