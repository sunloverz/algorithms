package queue

type Queue struct {
	elements []interface{}
	size int
}

func (queue *Queue) Enqueue(value interface{}) {
	queue.grow()
	queue.elements[queue.size] = value
	queue.size++
}

func (queue *Queue) Dequeue() interface{} {
	if queue.size == 0 {
		return nil
	}
	value := queue.elements[0]
	queue.elements = queue.elements[1:]
	queue.size--
	return value
}


func (queue *Queue) Peek() interface{}{
	 if cap(queue.elements) == 0 {
		return nil
	 }
	return queue.elements[0]
}

func (queue *Queue) Size() int {
	return queue.size;
}

func (queue *Queue) Values() []interface{} {
	return queue.elements[0:queue.size]
}

func (queue *Queue) grow() {
	currentCapacity := cap(queue.elements)
	if queue.size+1 > currentCapacity {
		queue.resize((currentCapacity + 1) * 2)
	}
}

func (queue *Queue) resize(length int) {
	newElements := make([]interface{}, length)
	copy(newElements, queue.elements)
	queue.elements = newElements
}

