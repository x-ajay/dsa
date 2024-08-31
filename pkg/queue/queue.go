package queue

import (
	"fmt"
)

type Queue struct {
	items []int
}

func New() *Queue {
	return &Queue{
		items: make([]int, 0),
	}
}

func (q *Queue) Enqueue(data int) {
	q.items = append(q.items, data)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front
}

func (q *Queue) Front() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.items[0]
}

func (q *Queue) Rear() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.items[len(q.items)-1]
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) String() string {
	result := ""
	for i := 0; i < len(q.items); i++ {
		result += fmt.Sprintf("[%d]->", q.items[i])
	}
	result += "nil"
	return result
}
