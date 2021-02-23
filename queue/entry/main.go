package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-go1/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println("q = ", q)
	q.Pop()
	q.Pop()
	fmt.Println("q = ", q)
	fmt.Println(q.IsEmpty())
	q.Pop()
	fmt.Println(q.IsEmpty())
}
