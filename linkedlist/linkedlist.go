package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head *Node
}

func (l *List) Insert(value int) {
	temp := Node{}
	temp.Value = value
	temp.Next = l.Head
	l.Head = &temp
}

func (l *List) InsertAt(n, value int) {
	temp1 := Node{}
	temp1.Value = value
	temp1.Next = nil
	if n == 1 {
		temp1.Next = l.Head
		l.Head = &temp1
		return
	}
	temp2 := l.Head
	for i:=0; i<n-2; i++ {
		temp2 = temp2.Next
	}
	temp1.Next = temp2.Next
	temp2.Next = &temp1
}

func (l *List) DeleteAt(n int) {
	temp := l.Head
	if n == 1 {
		l.Head = temp.Next
		temp = nil
		return
	}
	for i:=0; i<n-2; i++ {
		temp = temp.Next
	}
	temp2 := temp.Next
	temp.Next = temp2.Next
	temp2 = nil
}

func (l *List) Print() {

	for n := l.Head; n != nil; n = n.Next {
		fmt.Printf("%d ", n.Value)
	}
	fmt.Println()
}

func main() {
	list := List{}
	// list.Insert(2)
	// list.Insert(3)
	// list.Insert(4)
	// list.Insert(5)
	list.InsertAt(1, 2)
	list.InsertAt(2, 3)
	list.InsertAt(1, 4)
	list.InsertAt(2, 5)
	// list.Print()
	list.DeleteAt(1)
	list.Print()
	list.DeleteAt(3)
	list.Print()
	list.DeleteAt(2)
	list.Print()
}
