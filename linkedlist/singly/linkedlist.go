package singly

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
	for i := 0; i < n-2; i++ {
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
	for i := 0; i < n-2; i++ {
		temp = temp.Next
	}
	temp2 := temp.Next
	temp.Next = temp2.Next
	temp2 = nil
}

func (l *List) Values() []int {
	var values []int
	for n := l.Head; n != nil; n = n.Next {
		values = append(values, n.Value)
	}
	return values
}

func (l *List) Reverse() {
	var prev, current, next *Node
	next = l.Head.Next

	for current = l.Head; current != nil; current = next {
		next = current.Next
		current.Next = prev
		prev = current
	}
	l.Head = prev
}

func (l *List) BatchInsert(values []int) {
	for _, val := range values {
		l.Insert(val)
	}
}

func (l *List) Print() {
	for n := l.Head; n != nil; n = n.Next {
		fmt.Printf("%d ", n.Value)
	}
	fmt.Println()
}

func (l *List) RecursiveReverse(node *Node) {
	if node.Next == nil {
		l.Head = node
		return
	}
	l.RecursiveReverse(node.Next)
	next := node.Next
	next.Next = node
	node.Next = nil
}

func ReversePrint(node *Node) {
	if node == nil {
		return
	}
	ReversePrint(node.Next)
	fmt.Printf("%d ", node.Value)
}
