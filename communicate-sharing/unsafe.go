package main

import (
	"fmt"
	"sync"
)

var count1 = 0

func incUnsafe(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		count1++
	}
}

func countUnsafe() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go incUnsafe(&wg)
	}
	wg.Wait()
	fmt.Println(count1)
}
