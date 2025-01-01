package Qlinkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Queue struct {
	Head *Node
	Tail *Node
	Size int
}

func (q *Queue) Enqueue(Data int) {
	newNode := &Node{Data: Data}
	if q.Tail == nil {
		q.Tail = newNode
		q.Head = newNode
	} else {
		q.Tail.Next = newNode
		q.Tail = newNode
	}
	q.Size++
}

func (q *Queue) Dequeue() (int, bool) {
	if q.Head == nil {
		return 0, false
	}

	value := q.Head.Data
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
	q.Size--
	return value, true
}

func (q *Queue) PrintQueue() {
	current := q.Head
	for current != nil {
		fmt.Printf("%v ", current.Data)
		current = current.Next
	}
	fmt.Println()
}
