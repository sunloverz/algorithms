package linkedlistqueue

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Head *Node
	Rear *Node
}

func (q *Queue) Enqueue(value int) {
	temp := Node{Value: value}
	if q.Head == nil && q.Rear == nil {
		q.Head, q.Rear = &temp, &temp
		return
	}
	q.Rear.Next = &temp
	q.Rear = &temp
}

func (q *Queue) Deque() {
	if q.IsEmpty() {
		return
	}
	if q.Head == q.Rear {
		q.Head = nil
		q.Rear = nil
		return
	}
	temp := q.Head
	q.Head = temp.Next
	temp = nil
}

func (q *Queue) Peek() int {
	return q.Head.Value
}

func (q *Queue) IsEmpty() bool {
	return q.Head == nil && q.Rear == nil
}

func (q *Queue) Values() []int {
	values := []int{}
	for n := q.Head; n != nil; n = n.Next {
		values = append(values, n.Value)
	}
	return values
}