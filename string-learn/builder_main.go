package main

import (
	"fmt"
	"strings"
)

func test() {
	fmt.Print("hello")
}

func main() {
	test()

	var strBuilder strings.Builder
	strBuilder.Write([]byte{1, 2})
	strBuilder.WriteByte(2)
	strBuilder.WriteRune([]rune("中")[0])
	strBuilder.WriteString("中国")

}
