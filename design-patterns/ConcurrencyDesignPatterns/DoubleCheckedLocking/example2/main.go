package main

import (
	"log"
	"sync"
)

/*

 */

var arrMu sync.Mutex
var arr []int

// getArr retrieves arr, lazily initializing if needed. Double-checked locking
// avoids locking the entire function, and ensures that arr will be
// initialized only once.
func getArr() *[]int {

	if arr != nil { // 1st check
		return &arr
	}

	arrMu.Lock()
	defer arrMu.Unlock()

	if arr != nil { // 2nd check
		return &arr
	}
	arr = []int{0, 1, 2}
	return &arr
}

func main() {
	// thanks to double-checked locking, two goroutines attempting to getArr()
	// will not cause double-initialization
	go getArr()
	go getArr()

	log.Printf("%+v", getArr())
}
