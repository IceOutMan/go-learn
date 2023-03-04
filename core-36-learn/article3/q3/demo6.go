package main

import (
	//"com.meiken/core-36-learn/article3/q3/lib/internal" // 非法引入
	"com.meiken/core-36-learn/article3/q3/lib"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
}
