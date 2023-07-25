package main

import "fmt"

func map_test() {

	myMap := map[string]int{
		"zs": 11,
		"ls": 22,
		"ww": 33,
	}

	key := "ls"
	if value, ok := myMap[key]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("not found key:", key)
	}
	// 删除
	delete(myMap, "zs")

	// 遍历
	for key, val := range myMap {
		fmt.Println("key:", key, " val:", val)
	}

}
