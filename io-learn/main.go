package main

import (
	"fmt"
	"os"
)

func main() {
	wordFile, err := os.Open("word.txt")
	if err != nil {
		return
	}
	defer wordFile.Close()

	myContent := make([]byte, 10)
	n, err := wordFile.Read(myContent)
	if err != nil {
		return
	}
	fmt.Println(n)
	fmt.Printf("%s", myContent)

	file, err := os.ReadFile("word.txt")
	if err != nil {
		return
	}
	fmt.Printf("%s", file)

}
