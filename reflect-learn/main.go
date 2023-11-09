package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	var age int = 10
	fmt.Println(age)
	ageType := reflect.TypeOf(age)
	fmt.Println(ageType.String())

	myCat := Cat{"xiaohua", 10}
	catType := reflect.TypeOf(myCat)
	fmt.Println(catType.String())

	fmt.Println("===============")

	var xCat interface{} = myCat
	fmt.Println(reflect.TypeOf(xCat))
	fmt.Println(reflect.ValueOf(xCat))
	fmt.Println(reflect.ValueOf(xCat).NumField())
	fmt.Println(reflect.ValueOf(xCat).Field(0).Type())
	fmt.Println(reflect.ValueOf(xCat).Field(1).Type())
	fmt.Println(reflect.ValueOf(xCat).FieldByName("Name").Type())

}
