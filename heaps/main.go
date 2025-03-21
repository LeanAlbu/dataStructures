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

func (h *heap) heapifyDown(index int) {  // Changed to take index as parameter
	length := len(h.data)               // Fixed typo 'lenght' to 'length'
	for {
		largest := index
		left := 2*index + 1
		right := 2*index + 2

		if left < length && h.data[left] > h.data[largest] {
			largest = left
		}
		if right < length && h.data[right] > h.data[largest] {
			largest = right
		}
		if largest != index {
			h.data[index], h.data[largest] = h.data[largest], h.data[index]
			index = largest
		} else {
			break
		}
	}
}

func (h *heap) remove() int {
	if len(h.data) == 0 {
		return -1
	}
	max := h.data[0]
	last := len(h.data) - 1           // Fixed syntax error in variable declaration
	h.data[0] = h.data[last]
	h.data = h.data[:last]

	if len(h.data) > 0 {              // Fixed typo 'data.h' to 'h.data'
		h.heapifyDown(0)              // Pass 0 as starting index
	}
	return max                        // Return the removed max value instead of 0
}

func main() {
	myHeap := newHeap()
	myHeap.insert(10)
	myHeap.insert(5)
	myHeap.insert(15)
	myHeap.insert(20)

	// Should output: [20 15 10 5]
	// This is a max heap where parent is always greater than children
	fmt.Println(myHeap.data)
}
