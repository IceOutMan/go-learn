package main

import "fmt"

type MyString1 string
type MyString2 string

func main() {
	var container = []string{"zs", "ls", "ww"}
	// 1. container ---> interface{}
	// 2. 判断是否是 []string 类型
	// 如果是true，被判断的值将会被自动转换为[]string类型的值，并赋给变量value，否则value将被赋予nil（即“空”）。
	value, ok := interface{}(container).([]string)
	fmt.Println(value, ok)

	var s1 MyString1
	var s2 MyString2
	//var sArray1 []MyString1
	//var sArray2 []MyString2

	s1 = "s1"
	s2 = "s2"
	//sArray1 = []MyString1{"hhh", "hhh"}
	//sArray2 = []MyString2{"hhh", "hhh"}

	s1To2 := MyString2(s1)

	fmt.Println(string(s1To2))
	fmt.Println(s2)
}
