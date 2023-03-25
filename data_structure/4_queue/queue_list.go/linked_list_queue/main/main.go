package main

type ArrayQueue struct {
	items []interface{}
	front int
	rear  int
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		items: make([]interface{}, 0),
		front: 0,
		rear:  0,
	}
}

func (q *ArrayQueue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
	q.rear++
}

func (q *ArrayQueue) Dequeue() interface{} {
	if q.front == q.rear {
		return nil
	}
	item := q.items[q.front]
	q.front++
	return item
}

func (q *ArrayQueue) Size() int {
	return q.rear - q.front
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.front == q.rear
}
