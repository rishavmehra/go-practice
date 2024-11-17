package main

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	Head *Node
	Size int
}

func (s *Stack) Push(Data int) {
	newNode := &Node{Data: Data}
	s.Head = newNode
	s.Size++
}

func (s *Stack) Pop(data int) (int, bool) {
	removed := s.Head
	s.Head = s.Head.Next
	s.Size--
	return removed.Data, true

}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return 0
	}
	return s.Head.Data

}

func (s *Stack) IsEmpty() bool {
	return s.Head == nil
}
