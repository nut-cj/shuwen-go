package main

import (
	"fmt"
)

func main() {
	fmt.Println("***switch***")
	fmt.Println("res = ", swit(5, 3, "+"))
	fmt.Println("res2 = ", swit(5, 3, "-"))
	fmt.Println("res3 = ", swit(5, 3, "*"))
	fmt.Println("res4 = ", swit(5, 3, "/"))
	fmt.Println("res5 = ", swit(5, 3, "%"))
	// fmt.Println("res6 = ", swit(5, 3, "!"))
}

func swit(i, j int, expression string) int {
	var res int
	switch expression {
	case "+":
		res = i + j
	case "-":
		res = i - j
	case "*":
		res = i * j
	case "/":
		res = i / j
	case "%":
		res = i % j
	default:
		panic("err:" + expression)
	}
	return res
}
