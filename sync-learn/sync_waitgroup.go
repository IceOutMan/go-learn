package main

import "sync"

/*
*
 */
func main() {
	var wg sync.WaitGroup
	wg.Add(20)
	wg.Wait() // 等待在此处
	wg.Done() // goroutine 中进行，wg中的数字就减去1
}
