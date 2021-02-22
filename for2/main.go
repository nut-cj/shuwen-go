package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("***for2***")
	readFile("name.txt")
}

func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
