package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	// 仅控制一部分代码的原子性
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atomicInt
	a.increment()

	go func() {
		a.increment()
	}()

	time.Sleep(2000)
	fmt.Println(a.get())
}