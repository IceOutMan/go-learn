package main

import (
	"container/list"
	"fmt"
)

func list_test() {
	// 声明时不初始化, 内部为 nil
	var list_1 = list.List{}
	fmt.Println(list_1, "    ", list_1.Len())
	// 声明时初始化 list -> 初始化一个元素为0的 List
	var list_2 = list.New()
	fmt.Println(list_2, "    ", list_2.Len())
}
