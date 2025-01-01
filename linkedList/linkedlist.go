package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (dl *DoubleLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if dl.head == nil {
		dl.head = newNode
		dl.tail = newNode
	} else {
		dl.tail.next = newNode
		newNode.prev = dl.tail
		dl.tail = newNode
	}
}

func (dl *DoubleLinkedList) DeleteNode(data int) {
	current := dl.head
	for current != nil {
		if current.data == data {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				dl.head = current.next
			}

			if current.next != nil {
				current.next.prev = current.prev
			} else {
				dl.tail = current.prev
			}
			break
		}
		current = current.next
	}
}

func (dl *DoubleLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	if dl.head == nil {
		dl.head = newNode
		dl.tail = newNode
	} else {
		newNode.next = dl.head
		dl.head.prev = newNode
		dl.head = newNode
	}
}

func (dl *DoubleLinkedList) TraverseForward() {
	current := dl.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (dl *DoubleLinkedList) TraverseBackward() {
	current := dl.tail
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.prev
	}
	fmt.Println()
}

func main() {
	dl := &DoubleLinkedList{}

	dl.InsertAtEnd(10)
	dl.InsertAtEnd(20)
	dl.InsertAtEnd(30)
	dl.InsertAtBeginning(5)
	dl.TraverseForward()

	fmt.Println()
	dl.DeleteNode(20)
	dl.TraverseForward()

}
