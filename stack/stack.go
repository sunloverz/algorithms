package stack

type Stack struct {
	elements []interface{}
	size     int
}

func (stack *Stack) Push(value interface{}) {
	stack.grow()
	stack.elements[stack.size] = value
	stack.size++
}

func (stack *Stack) Pop() interface{} {
	stack.size--
	value := stack.elements[stack.size]
	stack.elements[stack.size] = nil
	return value
}

func (stack *Stack) Peek() interface{} {
	return stack.elements[stack.size-1]
}

func (stack *Stack) Search(searchValue interface{}) int {
	for i, value := range stack.elements {
		if searchValue == value {
			return i
		}
	}
	return -1
}

func (stack *Stack) Size() interface{} {
	return stack.size
}

func (stack *Stack) Empty() bool {
	return stack.size == 0
}

func (stack *Stack) grow() {
	currentCapacity := cap(stack.elements)
	if stack.size+1 > currentCapacity {
		stack.resize((currentCapacity + 1) * 2)
	}
}

func (stack *Stack) resize(length int) {
	newElements := make([]interface{}, length, length)
	copy(newElements, stack.elements)
	stack.elements = newElements
}
