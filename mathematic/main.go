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
		//1=49,2=50,3=51,4=52,5=53,6=54,7=55,8=56,9=57
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	arr, _ := Random(strs, 27)
	arrByte := []byte(arr)
	arrByte[0], arrByte[2], arrByte[4] = 49, 49, 49
	if arr[0] != 49 || arr[2] != 49 || arr[4] != 49 {
		arrByte[0] = 49
		arrByte[2] = 49
		arrByte[4] = 49
	}

	for i := 0; i < 24; i++ {
		if arrByte[i] == 2 {
			if i < 18 {
				arrByte[i+3] = 2
				arrByte[i+6] = 2
			}
		}
	}

	fmt.Println(string(arr))
	fmt.Println(string(arrByte))
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
