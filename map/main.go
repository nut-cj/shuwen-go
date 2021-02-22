package main

import (
	"fmt"
)

func main() {
	fmt.Println("***map****")
	// mapString()
	// mapMake()
	// mapNil()
	mapDelete()
}

func mapString() {
	m := map[string]string{
		"name":   "shuwen",
		"course": "mathematics",
		"site":   "www.xstiku.com",
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println(m["course"])

	na, ok := m["name"]
	fmt.Println(na, ok)

	if si, ok := m["sites"]; ok {
		fmt.Println(si, ok)
	} else {
		fmt.Println("key doses not exist!")
	}

	sit, ok := m["site"]
	fmt.Println(sit, ok)
}

func mapMake() {
	m := make(map[string]int)
	fmt.Println("m = ", m) // m == empty
}

func mapNil() {
	var m map[string]int
	fmt.Println("m = ", m) // m == nil
}

func mapDelete() {
	m := map[string]string{
		"name":   "shuwen",
		"course": "mathematics",
		"site":   "www.xstiku.com",
	}
	course, ok := m["course"]
	fmt.Println(course, ok)
	delete(m, "course") // delete
	course, ok = m["course"]
	fmt.Println(course, ok)
}
