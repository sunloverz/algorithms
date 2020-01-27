package doubly

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type List struct {
	Head *Node
}

func (l *List) InsertAtHead(value int) {
	newNode := Node{Value: value}
	if l.Head == nil {
		l.Head = &newNode
		return
	}
	newNode.Next = l.Head
	l.Head.Prev = &newNode
	l.Head = &newNode
}

func (l *List) Values() []int {
	values := []int{}
	for n := l.Head; n != nil; n = n.Next {
		values = append(values, n.Value)
	}
	return values
}
