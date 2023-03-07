package main

import "fmt"

func main() {

	// range
	numbers := [...]int{1, 2, 3, 4, 5, 6}
	maxLen := len(numbers) -1
	for index,value := range  numbers{
		if index == maxLen {
			numbers[0] += value
		}else{
			numbers[index + 1] += value
		}
	}

	fmt.Println(numbers)

}
