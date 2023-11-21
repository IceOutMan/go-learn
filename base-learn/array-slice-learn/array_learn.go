package main

import "fmt"

func arr() {
	fmt.Println("================== ARRAY ==================")

	// len: 7 , cap: 7
	arr := [7]int{1, 2, 3, 4, 5}
	fmt.Printf("len: %d , cap: %d\n", len(arr), cap(arr))

	// len: 5 , cap: 5
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("len: %d , cap: %d\n", len(slice), cap(slice))

}
