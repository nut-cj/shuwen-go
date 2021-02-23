package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// arr := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9}
	arr := make([]int, 27)
	arr[0], arr[2], arr[4] = 1, 1, 1
	rand.Seed(time.Now().Unix())
	for k := range arr {
		res := randInt(2, 9)
		arr[k] = res
		if arr[0] != 1 || arr[2] != 1 || arr[4] != 1 {
			arr[0] = 1
			arr[2] = 1
			arr[4] = 1
		}

		for i := 0; i < 27; i++ {
			if arr[i] == 2 {
				if i < 21 {
					arr[i+3] = 2
					arr[i+6] = 2
				}
			}
		}

		for i := 0; i < 27; i++ {
			if arr[i] == 3 {
				if i < 19 {
					arr[i+4] = 3
					arr[i+8] = 3
				}
			}
		}

		for i := 0; i < 27; i++ {
			if arr[i] == 4 {
				if i < 17 {
					arr[i+5] = 4
					arr[i+10] = 4
				}
			}
		}

		for i := 0; i < 27; i++ {
			if arr[i] == 5 {
				if i < 15 {
					arr[i+6] = 5
					arr[i+12] = 5
				}
			}
		}

		for i := 0; i < 27; i++ {
			if arr[i] == 6 {
				if i < 13 {
					arr[i+7] = 6
					arr[i+14] = 6
				}
			}
		}

		for i := 0; i < 27; i++ {
			if arr[i] == 7 {
				if i < 11 {
					arr[i+8] = 7
					arr[i+16] = 7
				}
			}
		}

		for i := 0; i < 27; i++ {
			if arr[i] == 8 {
				if i < 9 {
					arr[i+9] = 8
					arr[i+18] = 8
				}
			}
		}

		for i := 0; i < 27; i++ {
			if arr[i] == 9 {
				if i < 7 {
					arr[i+10] = 9
					arr[i+20] = 9
				}
			}
		}
	}
	fmt.Println("arr = ", arr)
}

func randInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}
