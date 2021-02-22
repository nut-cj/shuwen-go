package main

import (
	"fmt"
)

func main() {
	arr()
}

func arr() {
	var arr [3]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{1, 3, 5, 7, 9}
	fmt.Println(arr, arr2, arr3)
}
