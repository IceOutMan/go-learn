package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Name string
}

func (m MyStruct) Hello() string {
	return "Hello, " + m.Name + "!"
}

func DynamicCallMethod() {
	obj := MyStruct{Name: "Alice"}

	// 获取对象的反射值
	value := reflect.ValueOf(obj)

	// 获取方法并调用
	method := value.MethodByName("Hello")
	// Hello 方法没有参数，因此传递了一个空的 []reflect.Value{} 切片
	result := method.Call([]reflect.Value{})

	// 提取返回值
	if len(result) > 0 {
		// Interface() 方法将 reflect.Value 转换为接口类型
		// 通过将 reflect.Value 转换为 interface{}，你可以获取底层值的原始类型，并进一步进行类型断言或其他操作。
		// reflect.Value，它并不直接暴露底层的原始值，而是提供了一系列反射方法（如 Int()、String() 等）来获取底层值的不同表示
		returnValue := result[0].Interface()
		fmt.Println(returnValue)
	}
}
