package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("***for***")
	fmt.Println(convert2Bin(8))
	fmt.Println(convert2Bin(12)) // 8 + 4
	fmt.Println(convert2Bin(13)) // 8 + 4 1101
}

func convert2Bin(i int) string {
	res := ""
	for ; i > 0; i /= 2 {
		b := i % 2 // 13%2=1 6%2=0 3%2=1 1%2=1 1011 => 1101
		res = strconv.Itoa(b) + res
	}
	return res
}
