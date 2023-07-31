package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	//var num2 int32
	//
	//for {
	//	if atomic.CompareAndSwapInt32(&num2, 10, 0) {
	//		fmt.Println("The second number has gone to zero.")
	//		break
	//	}
	//	time.Sleep(time.Millisecond * 500)
	//}

	// 示例-1
	var box atomic.Value

	fmt.Println("Copy box to box2.")
	box2 := box // 原子值在真正使用前可以被复制

	v1 := [...]int{1, 2, 3}
	fmt.Printf("Store %v to box.\n", v1)
	box.Store(v1)
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Printf("The value load from box2 is %v.\n", box2.Load())

}
