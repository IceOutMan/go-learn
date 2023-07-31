package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
	name     string
	category string
}

func (c Cat) Name() string {
	return c.name
}

func (c Cat) Category() string {
	return c.category
}

func (c *Cat) SetName(name string) {
	c.name = name
}

func pointer_type() {
	var cat *Cat
	var pet Pet = cat

	fmt.Println("==================================================")
	fmt.Println("cat :", cat, "cat is nil ? ", cat == nil)
	fmt.Println("pet : ", pet, "pet is nil ? ", pet == nil)

	fmt.Println("==================================================")
	fmt.Printf("The type of pet is %T \n", pet)
	fmt.Printf("The type of pet is %s \n", reflect.TypeOf(pet).String())
}
