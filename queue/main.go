package main

import Qlinkedlist "github.com/rishavmehra/golangPractice/queue/linkedlist"

func main() {
	que := &Qlinkedlist.Queue{}

	que.Enqueue(10)
	que.Enqueue(2)
	que.PrintQueue()
	que.Dequeue()

	que.PrintQueue()

}
