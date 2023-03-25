package main

type CircularQueue struct {
	data  []interface{}
	front int
	rear  int
	size  int
}

func NewCircularQueue(k int) *CircularQueue {
	return &CircularQueue{
		data:  make([]interface{}, k),
		front: 0,
		rear:  0,
		size:  0,
	}
}

func (q *CircularQueue) Enqueue(val interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.data[q.rear] = val
	q.rear = (q.rear + 1) % len(q.data)
	q.size++
	return true
}

func (q *CircularQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	val := q.data[q.front]
	q.front = (q.front + 1) % len(q.data)
	q.size--
	return val
}

func (q *CircularQueue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.data[q.front]
}

func (q *CircularQueue) Rear() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.data[(q.rear-1+len(q.data))%len(q.data)]
}

func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *CircularQueue) IsFull() bool {
	return q.size == len(q.data)
}
