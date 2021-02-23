package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	strs := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	a, _ := Random(strs, 27)
	fmt.Println(a)
}

func Random(strings []string, length int) (string, error) {
	if len(strings) <= 0 {
		return "", errors.New("not nil")
	}

	if length <= 0 || len(strings) <= length-1 {
		return "", errors.New("not more")
	}

	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := ""
	for i := 0; i < length; i++ {
		str += strings[i]
	}
	return str, nil
}
