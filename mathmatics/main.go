package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 1.随机生成1-9中的数，每个数出现3次，存入切片中
	// 2.对随机生成的27个数进行排序，1号是1
	fmt.Println("***mathematics***")
	for i := 0; i < 1; i++ {
		arr := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9}
		for k, v := range arr {
			if v == 1 {
				result := RandIntn(2, 9)
				arr[1] = result
				for i := 0; i < 27; i++ {
					if result == arr[k] {
						arr = append(arr[:k], arr[k+1:]...)
					}
				}
				arr[3] = result
				arr[4] = 1
				fmt.Println(arr)
			}
			if v == 2 {

				fmt.Println(arr)
			}
		}
	}
}

func RandIntn(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}
