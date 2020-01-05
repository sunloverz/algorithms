package arraylist

type ArrayList struct {
	elements []int
	size     int
}

func (list *ArrayList) Add(value int) {
	list.grow()
	list.elements[list.size] = value
	list.size++
}

func (list *ArrayList) Get(index int) int {
	return list.elements[index]
}

func (list *ArrayList) Set(index, value int) {
	list.elements[index] = value
}

func (list *ArrayList) Remove(index int) {
	// for i := index; i < list.size; i++ {
	// 	list.elements[i] = list.elements[i+1]
	// }
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
}

func (list *ArrayList) Values() []int {
	return list.elements[0:list.size]
}

func (list *ArrayList) Size() int {
	return list.size
}

func (list *ArrayList) Clear() {
	list.elements = []int{}
	list.size = 0
}

func (list *ArrayList) resize(cap int) {
	newElements := make([]int, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *ArrayList) grow(){
	currentCapacity := cap(list.elements)
	if list.size + 1 >= currentCapacity {
		newCapacity := (currentCapacity + 1) * 2
		list.resize(newCapacity)
	}
}