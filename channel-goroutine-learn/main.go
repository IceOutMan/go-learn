package main

import (
	"fmt"
	"time"
)

func main() {

	// var ch chan int ( 只声明，没有初始化，是一个nil，对 nil 的channel 都会被阻塞）

	//ch := make(chan int,3)

	//ch<- 1
	//ch <- 2
	//ch <- 3

	//element := <- ch
	//fmt.Println(element)
	//
	//element = <- ch
	//fmt.Println(element)
	//
	//element = <- ch
	//fmt.Println(element)

	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(time.Millisecond * 500)

}
