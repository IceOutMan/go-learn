package main

import "fmt"

type Dog struct {
	name string
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func (d *Dog) Name() string {
	return d.name
}

func main() {
	dog := Dog{name: "litter pig"}
	fmt.Println("dog name is: ", dog.Name())

	dog.SetName("monster")
	fmt.Println("dog after modify name is: ", dog.Name())
}
