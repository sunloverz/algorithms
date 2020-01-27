package linkedliststack

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Head *Node
}

func (stack *Stack) Push(value int) {
	newNode := Node{Value: value}
	newNode.Next = stack.Head
	stack.Head = &newNode
}

func (stack *Stack) Pop() int {
	node := stack.Head
	value := node.Value
	stack.Head = node.Next
	node = nil
	return value
}

func (stack *Stack) Peek() int {
	return stack.Head.Value
}

func (stack *Stack) IsEmpty() bool {
	return stack.Head == nil
}

func (stack *Stack) Values() []int {
	values := []int{}
	for n := stack.Head; n != nil; n = n.Next {
		values = append(values, n.Value)
	}
	return values
}
