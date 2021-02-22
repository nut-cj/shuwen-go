package main

import "fmt"

func main() {
	arr()
}

func arr() {
	var arr [3]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{1, 3, 5, 7, 9}
	fmt.Println(arr, arr2, arr3)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	for k, v := range arr2 {
		fmt.Println(k, v)
	}

	for k := range arr2 {
		fmt.Println(k)
	}
}
