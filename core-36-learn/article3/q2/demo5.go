package main

import (
	"com.meiken/core-36-learn/article3/q2/lib"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib5.Hello("")
}
