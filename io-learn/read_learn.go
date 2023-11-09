package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func read_line() {

	wordFile, err := os.Open("word.txt")
	if err != nil {
		return
	}
	defer wordFile.Close()

	scanner := bufio.NewScanner(wordFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("read err")
	}
}

func read_block() {

	wordFile, err := os.Open("word.txt")
	if err != nil {
		return
	}
	defer wordFile.Close()

	byteSlice := make([]byte, 1)

	for {
		n, err := wordFile.Read(byteSlice)
		if err != nil {
			return
		}
		if n == 0 || err == io.EOF {
			break
		}
		fmt.Printf("%d, %s \n", n, byteSlice)
	}
}

func random_read() {

	wordFile, err := os.Open("word.txt")
	if err != nil {
		return
	}
	defer wordFile.Close()

	fileReader := bufio.NewReader(wordFile)
	// peek到指定字节数的数据，当前指针不变
	var n int = 2
	byteSlice, err := fileReader.Peek(n)
	fmt.Printf("Peeked at %d bytes：%s\n", n, byteSlice)

	// 由于 peek 指针没有变化，此处读到 byteSlice 和上面一致
	numBytesRead, err := fileReader.Read(byteSlice)
	fmt.Printf("Read %d bytes:%s\n", numBytesRead, byteSlice)

}
