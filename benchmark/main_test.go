package main

import (
	"math/rand"
	"testing"
)

func sumElements(slice []int) int {
	sum := 0
	var tempSlice []int
	tempSlice = append(tempSlice, slice...)
	for _, v := range tempSlice {
		sum += v
	}
	return sum
}

func BenchmarkSumElements(b *testing.B) {
	b.ReportAllocs()
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = rand.Int()
	}

	// Reset the timer to exclude the time taken to
	// allocate the slice.
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sumElements(slice)
	}
}
