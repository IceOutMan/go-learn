package main

import (
	"fmt"
	"time"
)

func run_goroutine() {
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(time.Millisecond * 500)
}
