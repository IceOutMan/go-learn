package main

import (
	"fmt"
)

func main() {
	// arr()
	// slice()
	a := make([]int, 32)
	b := a[1:16]
	fmt.Printf("a=> len:%d,cap:%d\n", len(a), cap(a))
	fmt.Printf("b=> len:%d,cap:%d\n", len(b), cap(b))

	// append num
	fmt.Println("========> append num to a")
	a = append(a, 1)
	fmt.Printf("a=> len:%d,cap:%d\n", len(a), cap(a))
	fmt.Printf("b=> len:%d,cap:%d\n", len(b), cap(b))


	// var i reflect.Value
	// fmt.Println(i.Interface())

}
