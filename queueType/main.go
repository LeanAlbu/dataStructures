package main

import "fmt"

type queue struct {
	elements []string
	max      int
}

func main() {
	q := queue{max: 5}
	q.enqueue("JHON")
	q.enqueue("PEDRO")
	q.enqueue("JUAN")
	q.enqueue("CARLOS")
	q.enqueue("MIGUEL")

	q.printQueue()

	fmt.Println("Dequeue: ", q.dequeue())

}

func (q *queue) getSize() int {
	return len(q.elements)
}

func (q *queue) isEmpty() bool {
	return len(q.elements) == 0
}

func (q *queue) enqueue(element string) {
	if q.getSize() == q.max {
		fmt.Println("Queue is full")
		return
	} else {
		q.elements = append(q.elements, element)
	}
}

func (q *queue) dequeue() string {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return ""
	} else {
		element := q.elements[0]
		q.elements = q.elements[1:]
		return element
	}
}

func (q *queue) printQueue() {
	for _, element := range q.elements {
		fmt.Println(element)
	}
}
