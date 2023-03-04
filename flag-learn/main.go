package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var age int
var address = flag.String("address", "default address","this is address")


func init(){
	fmt.Println(os.Args[1:])

	flag.StringVar(&name, "name", "default name", "This is name")
	flag.IntVar(&age, "age", 12, "This is age")

	flag.Parse()
}

func main() {
	fmt.Printf("name is %s\n", name)
	fmt.Printf("age is  %d\n", age)
	fmt.Printf("address is  %s\n", *address)
}
