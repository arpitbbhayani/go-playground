package main

import (
	"fmt"
	"sync"
)

var count3 = 0
var ich = make(chan int)

func incPassing(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ich <- 1
	}
}

func consumer(mg *sync.WaitGroup) {
	defer mg.Done()
	for delta := range ich {
		count3 += delta
	}
}

func producer(mg *sync.WaitGroup) {
	defer mg.Done()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go incPassing(&wg)
	}
	wg.Wait()
	close(ich)
}

func countPassing() {
	var mg sync.WaitGroup

	mg.Add(2)
	go consumer(&mg)
	go producer(&mg)
	mg.Wait()

	fmt.Println(count3)
}
