package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func cas_test() {
	var num2 int32

	for {
		if atomic.CompareAndSwapInt32(&num2, 10, 0) {
			fmt.Println("The second number has gone to zero.")
			break
		}
		fmt.Println("Will Sleep Wait")
		time.Sleep(time.Millisecond * 500)
	}
}
