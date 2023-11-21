package main

import "fmt"

func main() {
	//arr := []int{1, 3, 5, 7}
	//
	//for i, val := range arr {
	//	fmt.Println(i, val)
	//}
	//
	//for i := range arr {
	//	fmt.Println(i)
	//}

	numbers2 := []int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, val := range numbers2 {
		fmt.Println(i, val)
		fmt.Println(numbers2)
		if i == maxIndex2 {
			numbers2[0] += val
		} else {
			numbers2[i+1] += val
		}
	}
	fmt.Println("Final :", numbers2)
}
