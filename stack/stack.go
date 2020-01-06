package stack

type Stack struct {
	elements []interface{}
	size int
}

func (stack *Stack) Push(value interface{}) {
	stack.elements[stack.size] = value
	stack.size++
}

func (stack *Stack) Pop() interface{} {
	value := stack.elements[stack.size]
	stack.size--
	return value
}

