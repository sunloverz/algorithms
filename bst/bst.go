package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (t *BinarySearchTree) Insert(value int) {
	t.Root = insertNode(t.Root, value)
}

func (t *BinarySearchTree) Search(value int) bool {
	return search(t.Root, value)
}

func search(root *Node, value int) bool {
	if root == nil {
		return false
	} else if root.Value == value {
		return true
	} else if value > root.Value {
		return search(root.Right, value)
	} else {
		return search(root.Left, value)
	}
}

func insertNode(root *Node, value int) *Node {
	if root == nil {
		root = &Node{Value: value}
	} else if value > root.Value {
		root.Right = insertNode(root.Right, value)
	} else {
		root.Left = insertNode(root.Left, value)
	}
	return root
}

func (t *BinarySearchTree) Min() int {
	var min int
	for n := t.Root; n != nil; n = n.Left {
		min = n.Value
	}
	return min
}

func (t *BinarySearchTree) Max() int {
	var max int
	for n := t.Root; n != nil; n = n.Right {
		max = n.Value
	}
	return max
}

func main() {
	var tree BinarySearchTree
	tree.Insert(15)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(6)
	tree.Insert(12)
	fmt.Println(tree.Root.Value)
	fmt.Println(tree.Search(6))
	fmt.Println(tree.Search(10))
	fmt.Println(tree.Search(25))
	fmt.Println(tree.Search(30))
	fmt.Println(tree.Min())
	fmt.Println(tree.Max())
}
