package main

import (
	"container/ring"
	"fmt"
)

func ring_test() {
	var my_ring = ring.New(5)

	for i := 1; i <= 5; i++ {
		my_ring.Value = i
		my_ring = my_ring.Next()
	}

	// 1 + 2 + ... + 5
	s := 0
	my_ring.Do(func(val interface{}) {
		s = s + val.(int)
	})

	fmt.Println(s)

}
