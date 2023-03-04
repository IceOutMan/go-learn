package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	// list
	var l = list.List{}
	fmt.Println(l,"    ",l.Len())


	// ring
	var r = ring.Ring{}
	fmt.Println(r, "    ", r.Len())

	// heap

	// map
	myMap := map[string]int{
		"zs" : 11,
		"ls" : 22,
		"ww" : 33,
	}

	key := "ls"
	if value, ok := myMap[key]; ok{
		fmt.Println(value)
	}else{
		fmt.Println("not found key:", key)
	}


}
