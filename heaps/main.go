package main

import "fmt"

type heap struct {
	data []int
}

func newHeap() *heap {
	return &heap{}
}

func (h *heap) heapifyUp() {
	index := len(h.data) - 1
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[index] > h.data[parent] {
			h.data[index], h.data[parent] = h.data[parent], h.data[index]
			index = parent
		} else {
			break
		}
	}
}	

func (h *heap) insert(data int) {
	h.data = append(h.data, data)
	h.heapifyUp()
}

func (h *heap) heapifyDown () {
	index := len(h.data)-1
	lenght := len(h.data)
	for{
		largest := index
		left := 2*index + 1
		right := 2*index + 2

		if left < lenght && h.data[left] > h.data[largest]{
			largest = left
		}
		if right < lenght && h.data[right] > h.data[largest]{
			largest = right
		}
		if largest != index{
			h.data[index], h.data[largest] = h.data[largest], h.data[index]
			index = largest
		}else{
			break
		}
	}
}

func (h *heap) remove() int {
	if len(h.data) == 0 {
		return -1
	}
	max := h.data[0]
	last :- len(h.data)-1
	h.data[0] = h.data[last]
	h.data = h.data[:last]

	if len(data.h) > 0{
		h.heapifyDown(0)
	}
	return 0
}












func main() {
	myHeap := newHeap()
	myHeap.insert(10)
	myHeap.insert(5)
	myHeap.insert(15)
	myHeap.insert(20)

	// Output: [15 5 10]
	for _, v := range myHeap.data {
		fmt.Println(v)
	}

}
