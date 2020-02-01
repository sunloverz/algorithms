package main

import (
	"algorithms/bst"
	"fmt"
)

func main() {
	var tree bst.BinarySearchTree
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
	fmt.Println("bfs travers tree")
	fmt.Println(tree.BFS())
}
