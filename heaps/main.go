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

func (h *heap) heapifyDown() {
	index := 0
	lastIndex := len(h.data) - 1
	for index < lastIndex {
		left := 2*index+1
		right := 2*index+2
		largest := index
		if left <= lastIndex && h.data[left] > h.data[largest] {
			largest = left
		}else{
			largest = index
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
