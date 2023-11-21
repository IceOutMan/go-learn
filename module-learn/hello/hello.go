package main

import (
	"com.meiken/module-learn/greeting"
	"fmt"
)

func main() {

	message := greeting.Hello("Gladys")
	fmt.Println(message)
}
