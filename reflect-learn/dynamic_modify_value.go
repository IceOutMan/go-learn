package main

import (
	"fmt"
	"reflect"
)

func DynamicModifyValue() {
	var x int = 42
	// reflect.ValueOf(&x) 获取了 x 的地址的反射值
	// 获取变量的反射值, 通过 .Elem() 方法获取到了指针指向的值的反射值
	value := reflect.ValueOf(&x).Elem()
	// 修改变量的值
	value.SetInt(100)
	fmt.Println(x) // 输出: 100
}
