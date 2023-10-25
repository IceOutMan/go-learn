package main

import "fmt"

func slice() {
	fmt.Println("================== SLICE ==================")

	// len: 5 , cap: 9
	make_slice1 := make([]int, 5, 9)
	fmt.Printf("len: %d , cap: %d\n", len(make_slice1), cap(make_slice1))

	// len: 5 , cap: 5
	make_slice2 := make([]int, 5)
	fmt.Printf("len: %d , cap: %d\n", len(make_slice2), cap(make_slice2))

	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice1 := arr2[3:6]

	// slice1 [4,5,6]
	// len: 3 , cap: 5
	fmt.Printf("len: %d , cap: %d\n", len(slice1), cap(slice1))

	arr3 := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "g", "k", "l"}
	str_slice_1 := arr3[2:9]

	// str_slice_1 [c, d, e, f, g , h, i]
	// len: 7 , cap : 10
	fmt.Printf("len: %d , cap: %d\n", len(str_slice_1), cap(str_slice_1))

	str_slice_2 := str_slice_1[3:5]

	// str_slice_2 [f, g]
	// len: 2 , cap : 7
	fmt.Printf("len: %d , cap: %d\n", len(str_slice_2), cap(str_slice_2))

	slice4 := make([]int, 3)
	// [0 , 0, 0]  len: 3 cap:3
	fmt.Printf("len: %d , cap: %d\n", len(slice4), cap(slice4))

	slice5 := append(slice4, 1)
	// [0,0,0,1] len: 4 cap:6
	fmt.Printf("len: %d , cap: %d\n", len(slice5), cap(slice5))

}
