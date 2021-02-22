package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("if")
	ifBranch()
}

func ifBranch() {
	const fileName = "helloName.txt"
	if contents, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
