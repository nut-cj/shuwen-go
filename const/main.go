package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("const")
	consts()
}

func consts() {
	const (
		fileName = "hello.txt"
		i, j     = 3, 4
	)
	var k int
	k = int(math.Sqrt(i*i + j*j))
	fmt.Println("k = ", k)
	fmt.Println("fileName = ", fileName)
}
