package main

import (
	"algorithms/linkedlist"
	"fmt"
)

func main() {
	list := linkedlist.List{}
	list.BatchInsert([]int{4, 3, 2, 1}) // 1 2 3 4
	list.RecursiveReverse(list.Head)
	fmt.Println(list.Values())
}
