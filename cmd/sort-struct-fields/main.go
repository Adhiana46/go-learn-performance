package main

import (
	"fmt"
	"unsafe"
)

type unsortedStruct struct {
	testBool1  bool    // 1 byte
	testFloat1 float64 // 8 bytes
	testBool2  bool    // 1 byte
	testFloat2 float64 // 8 bytes
}

type sortedStruct struct {
	testFloat1 float64 // 8 bytes
	testFloat2 float64 // 8 bytes
	testBool1  bool    // 1 byte
	testBool2  bool    // 1 byte
}

func main() {
	a := unsortedStruct{}
	fmt.Println("Unsorted Struct: ", unsafe.Sizeof(a)) // 32 bytes

	b := sortedStruct{}
	fmt.Println("Sorted Struct: ", unsafe.Sizeof(b))
}
