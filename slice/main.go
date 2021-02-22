package main

import (
	"fmt"
)

func main() {
	sliceDemo()
}

func sliceDemo() {
	arr := [...]int{1, 3, 5, 7, 9}
	fmt.Println("arr[2,4]=", arr[2:4]) // [2,4)=>[5,7]
	fmt.Println("arr[:4]=", arr[:4])   // [0,3]=>[1,3,5,7]
	fmt.Println("arr[2:]=", arr[2:])   // [2,len(s)]=>[5,7,9]
	fmt.Println("arr[:]=", arr[:])     // [2,4)=>[1,3,5,7,9]
}
