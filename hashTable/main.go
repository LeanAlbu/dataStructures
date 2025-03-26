package main

import (
	"fmt"
)

type Node struct {
	key   string
	value int
	next  *Node
}

type Hashtable struct {
	table []*Node
	size  int
}

func hashFunc(key string, size int) int {
	sum := 0

	for _, char := range key {
		sum += int(char)
	}
	return sum % size
}

func hashNew(size int) *Hashtable {
	return &Hashtable{
		table: make([]*Node, size),
		size:  size,
	}
}

func (h *Hashtable) hashInsert(key string, value int) {
	index := hashFunc(key, h.size)

	if h.table[index] == nil {
		h.table[index] = &Node{key: key, value: value}
		return
	}

	aux := h.table[index]
	for aux.next != nil {
		if aux.key == key {
			aux.value = value
			return
		}
		aux = aux.next
	}
	if aux.key == key {
		aux.value = value
	} else {
		aux.next = &Node{key: key, value: value}
	}
}

func (h *Hashtable) hashSearch(key string) (int, bool) {
	index := hashFunc(key, h.size)
	aux := h.table[index]
	for aux != nil {
		if aux.key == key {
			return aux.value, true
		}
		aux = aux.next
	}
	return 0, false
}

func (h *Hashtable) hashRemove(key string) (int, bool) {
	index := hashFunc(key, h.size)
	aux := h.table[index]
	if aux != nil && aux.key == key {
		removed := h.table[index].value
		h.table[index] = aux.next
		return removed, true
	}
	for aux != nil && aux.next != nil {
		if aux.next.key == key {
			removed := aux.next.value
			aux.next = aux.next.next
			return removed, true
		}
		aux = aux.next
	}
	return 0, false
}

func main() {

}
