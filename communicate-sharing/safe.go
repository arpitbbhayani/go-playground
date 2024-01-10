package main

import (
	"fmt"
	"sync"
)

var count2 = 0
var mu sync.Mutex

func incSafe(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mu.Lock()
		count2++
		mu.Unlock()
	}
}

func countSafe() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go incSafe(&wg)
	}
	wg.Wait()
	fmt.Println(count2)
}
