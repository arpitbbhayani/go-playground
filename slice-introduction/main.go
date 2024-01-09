package main

import "fmt"

func printLines() {
	fmt.Println("--------------------------------------------------")
}

func printArray(name string, arr [10]int) {
	fmt.Printf("%s: %T %v %p %p\n", name, arr, arr, &arr, &arr[0])
}

func printSlice(name string, s []int) {
	fmt.Printf("%s: %T %v %p %p\n", name, s, s, &s, &s[0])
}

func main() {
	// defining the array
	var arr [10]int = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	printArray("arr", arr)
	printLines()

	// creating the slice from the entire array
	var s1 = arr[:]
	printSlice("s1", s1)
	printLines()

	var s2 = arr[4:6]
	printSlice("s2", s2)
	printLines()

	// creating a new slice from another slice
	var s3 = s1
	printSlice("s3", s3)
	printLines()

	// updating a slice, seeing the changes reflect in another slice and base array
	s3[0] = 123
	printSlice("s3", s3)
	printSlice("s2", s2)
	printSlice("s1", s1)
	printArray("arr", arr)
	printLines()

	// when we append
	var s4 = append(s2, 101, 102, 103)
	printSlice("s4", s4)
	printSlice("s3", s3)
	printSlice("s2", s2)
	printSlice("s1", s1)
	printArray("arr", arr)
	printLines()

	// when we append more than the size of the underlying array
	var s5 = append(s2, 201, 202, 203, 204, 205, 206)
	printSlice("s5", s5)
	printSlice("s4", s4)
	printSlice("s3", s3)
	printSlice("s2", s2)
	printSlice("s1", s1)
	printArray("arr", arr)
	printLines()
}
