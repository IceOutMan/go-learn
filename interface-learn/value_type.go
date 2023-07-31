package main

import (
	"fmt"
	"reflect"
)

type Dog struct {
	name     string
	category string
}

func (d Dog) Name() string {
	return d.name
}

func (d Dog) Category() string {
	return d.category
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func value_type() {
	dog := Dog{name: "litter pig"}
	var pet Pet = dog

	dog.SetName("monster")
	fmt.Println("dog Name: ", dog.Name())
	fmt.Println("pet Name: ", pet.Name())

	fmt.Println("==================================================")

	fmt.Printf("The type of pet is %T \n", pet)
	fmt.Printf("The type of pet is %s \n", reflect.TypeOf(pet).String())
}
