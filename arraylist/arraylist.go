package arraylist

type ArrayList struct {
	elements []interface{}
	size     int
}

func (list *ArrayList) Add(values ...interface{}) {
	list.grow(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

func (list *ArrayList) Get(index int) interface{} {
	return list.elements[index]
}

func (list *ArrayList) Set(index, value int) {
	list.elements[index] = value
}

func (list *ArrayList) Remove(index int) {
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
}

func (list *ArrayList) Values() []interface{} {
	newElements := make([]interface{}, list.size)
	copy(newElements, list.elements)
	return newElements
}

func (list *ArrayList) Size() int {
	return list.size
}

func (list *ArrayList) Clear() {
	list.elements = []interface{}{}
	list.size = 0
}

func (list *ArrayList) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *ArrayList) grow(length int) {
	currentCapacity := cap(list.elements)
	if list.size+length >= currentCapacity {
		newCapacity := (currentCapacity + length) * 2
		list.resize(newCapacity)
	}
}
