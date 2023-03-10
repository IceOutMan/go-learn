package main

import (
	"sync"
)

func main() {
	var sMap sync.Map

	sMap.Store("kk", "tt")
}
