package main

import (
	"fmt"
)

type treeNode struct {
	val         int
	left, right *treeNode
}

func main() {
	fmt.Println("***struct***")
	tree()
}

func createNode(value int) *treeNode {
	return &treeNode{val: value}
}

func (node treeNode) print() {
	fmt.Println(node.val)
}

func (node *treeNode) setValue(value int) {
	node.val = value
}

func tree() {
	var root treeNode
	root = treeNode{val: 5}
	fmt.Println("root = ", root)

	root2 := treeNode{
		val: 3,
	}
	fmt.Println("root = ", root2)

	root.left = &treeNode{}
	root.right = &treeNode{7, nil, nil}
	fmt.Println("root = ", root)
	root.right.left = new(treeNode)
	nodes := []treeNode{
		{val: 3},
		{},
		{9, nil, &root},
	}
	fmt.Println("nodes = ", nodes)
	root.left.right = createNode(30)
	fmt.Println("nodes = ", nodes)
	root.print()

	root.right.left.setValue(90)
	root.right.left.print()
}
