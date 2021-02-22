package main

import (
	"fmt"
)

func main() {
	fmt.Println("***switch2***")
	// level(101)
	fmt.Println("leve=", level(99))
	fmt.Println("leve2=", level(89))
	fmt.Println("leve3=", level(69))
}

func level(score int) string {
	lev := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Error score:%d", score))
	case score < 70:
		lev = "C"
	case score < 90:
		lev = "B"
	case score <= 100:
		lev = "A"
	}
	return lev
}
