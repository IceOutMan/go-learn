package main

import (
	"fmt"
	"sync"
	"time"
)

type book struct{
	content string
	lock sync.RWMutex
}

func (b *book) read() string{
	b.lock.RLock()
	defer b.lock.RUnlock()

	return b.content
}

func (b *book) write(){
	b.lock.Lock()
	defer b.lock.Unlock()
	time.Sleep(2000)
	b.content = b.content + "HH_"
}

func main() {
	var b book

	b.write()

	go func() {
		for i:=0;i<10;i++ {
			fmt.Println(b.read())
		}
	}()

	time.Sleep(5000)
	fmt.Println(b.content)

}
