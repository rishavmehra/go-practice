package QSlice

type Queue struct {
	item []int
}

func (q *Queue) Enqueue(additem int) {
	q.item = append(q.item, additem)
}

func (q *Queue) Dequeue() int {
	ritem := q.item[0]
	q.item = q.item[1:]
	return ritem
}

func (q *Queue) IsEmpty() bool {
	return len(q.item) == 0
}
