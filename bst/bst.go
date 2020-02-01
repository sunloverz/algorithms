package bst

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

func (t *BinarySearchTree) BFS() []int {
	if t.Root == nil {
		return []int{}
	}
	values := make([]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, t.Root)
	for ok := true; ok; ok = len(queue) > 0 {
		current := queue[0]
		values = append(values, current.Value)
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
		queue = queue[1:]
	}
	return values
}