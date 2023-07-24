package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	var container = []string{"zs", "ls", "ww"}
	// 1. container ---> interface{}
	// 2. 判断是否是 []string 类型
	// 如果是true，被判断的值将会被自动转换为[]string类型的值，并赋给变量value，否则value将被赋予nil（即“空”）。
	value, ok := interface{}(container).([]string)
	fmt.Println(value, ok)
	i := make([]interface{}, 3) //定义一个长度为3的空接口数组
	i[0] = 1
	i[1] = "go"
	i[2] = Student{"loria", 25}
	//类型查询,类型断言
	for index, data := range i {
		//第一个返回的是值,第二个返回的是判断内容的真假
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int,内容为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为string,内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为Student,内容为name=%s,age=%d\n", index, value.name, value.age)
		}
	}
}
