package main

import (
	"fmt"
)

func main() {
	fmt.Println("enums")
	enums()
	enums2()
}

func enums() {
	const (
		Shuwen = iota
		_
		Richard
		Ritchie
		Sophie
	)
	fmt.Println(Shuwen, Sophie, Richard, Ritchie)
}

func enums2() {
	const ( // b,kb,mb,gb,tb,pb
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
