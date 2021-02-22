package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	fmt.Println("invar")
	cmplx2()
	triangle()
}

func cmplx2() {
	c := 3 + 4i // 3^2+4^2=5^2
	fmt.Println(cmplx.Abs(c))
}

func triangle() {
	var i, j int = 3, 4
	var k int
	k = int(math.Sqrt(float64(i*i + j*j)))
	fmt.Println("k = ", k)
}
