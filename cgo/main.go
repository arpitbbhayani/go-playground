package main

/*
#include "add.c"
*/
import "C"

import "fmt"

func main() {
	result := C.add(5, 3)
	fmt.Println("sum:", result)
}
