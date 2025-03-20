package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	size int
}

func createNode(data int) *node {
	return &node{data: data, next: nil}
}

func llcreate() *linkedList {
	return &linkedList{head: nil, size: 0}
}

func (ll *linkedList) llappend(data int) {
	newNode := createNode(data)
	if ll.head == nil {
		ll.head = newNode
	} else {
		aux := ll.head
		for aux.next != nil {
			aux = aux.next
		}
		aux.next = newNode
	}
	ll.size++
}

func (ll *linkedList) llprepend(data int) {
	newNode := createNode(data)
	if ll.head == nil {
		ll.head = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
	ll.size++
}

func (ll *linkedList) lldelete(data int) {
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	} else {
		if ll.head.data == data {
			ll.head = ll.head.next
			ll.size--
		} else {
			aux := ll.head
			for aux.next != nil {
				if aux.next.data == data {
					aux.next = aux.next.next
					ll.size--
					return
				}
			}
		}
	}
}

func (ll *linkedList) llprint() {
	if ll.head == nil {
		fmt.Println("List is empty")
	} else {
		aux := ll.head
		for aux != nil {
			fmt.Print(aux.data)
			aux = aux.next
		}
	}
}
