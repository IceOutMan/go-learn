package main

import (
	"fmt"
	"sync/atomic"
)

func atomic_add() {

	var age int32 = 2

	fmt.Println("init age: ", age)
	atomic.AddInt32(&age, 10)
	fmt.Println("after atomic.AddInt32(&age,10) ->  age: ", age)

}
