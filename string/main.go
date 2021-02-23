package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("***string***")
	stringByte()
}

func stringByte() {
	s := "hello学思题库!" // UTF-8
	fmt.Println(s)
	for _, v := range []byte(s) {
		fmt.Printf("%X ", v)
	}
	fmt.Println()

	for k, v := range s {
		fmt.Printf("[%d %X],", k, v)
	}
	fmt.Println()
	fmt.Println("Rune count = ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	fmt.Println()

	for k, v := range []rune(s) {
		fmt.Printf("[%d %c]", k, v)
	}
	fmt.Println()
}
