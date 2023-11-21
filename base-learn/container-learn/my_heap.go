package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	// 递增 -> 也就是小顶堆
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}

// heap.Pop()
//  1. 会交换堆顶元素到末尾 swap(start,end)
//  2. 堆化 down(start,end-1)
//
// 所以我们的实现直接从末尾 pop
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(*h)
	val := old[n-1]

	*h = old[:n-1]
	return val
}

func heap_test() {

	var my_heap = &IntHeap{5, 3, 8, 9, 6}
	heap.Init(my_heap)

	heap.Push(my_heap, 18)

	for len(*my_heap) > 0 {
		fmt.Print(heap.Pop(my_heap), "  ")
	}
	fmt.Println()
}
