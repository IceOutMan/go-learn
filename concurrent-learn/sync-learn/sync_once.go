package main

import (
	"fmt"
	"sync"
)

/*
*
 */
func main() {
	var oc sync.Once

	oc.Do(func() {
		fmt.Println("oc done one")
	})

	oc.Do(func() {
		fmt.Println("oc done two")
	})
}
