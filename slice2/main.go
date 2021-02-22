package main

import (
	"fmt"
)

func main() {
	fmt.Println("***append***")
	// sliceAppend()
	// copySlice()
	// deleteSlice()
	// popFrontSlice()
	popTailSlice()
}

func printSlice(s []int) {
	fmt.Printf("%v,len = %d,cap = %d\n", s, len(s), cap(s))
}

func sliceAppend() {
	var s []int
	count := 100
	for i := 0; i < count; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s2 := []int{1, 3, 5}
	printSlice(s2)

	s3 := make([]int, 10)
	printSlice(s3)

	s4 := make([]int, 10, 32)
	printSlice(s4)

	copy(s2, s3)
	printSlice(s2)
}

func copySlice() {
	s := []int{1, 3, 5}
	s2 := []int{2, 4, 6, 8}
	copy(s, s2)
	fmt.Println(s)
}

func deleteSlice() {
	s := []int{1, 3, 5, 7, 9}
	s = append(s[:2], s[3:]...)
	fmt.Println(s)
}

func popFrontSlice() {
	s := []int{1, 3, 5, 7, 9}
	front := s[0]
	s = s[1:]
	fmt.Println(front)
	fmt.Println(s) // 3,5,7,9
}

func popTailSlice() {
	s := []int{1, 3, 5, 7, 9}
	tail := s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Println(tail) // 9
	fmt.Println(s)    // 1,3,5,7
}
