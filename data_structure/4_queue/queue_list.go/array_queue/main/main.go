package queue

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

type LinkedListQueue struct {
	front *ListNode
	rear  *ListNode
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		front: nil,
		rear:  nil,
	}
}

func (q *LinkedListQueue) Enqueue(item interface{}) {
	node := &ListNode{Val: item, Next: nil}
	if q.rear == nil {
		q.front = node
		q.rear = node
	} else {
		q.rear.Next = node
		q.rear = q.rear.Next
	}
}

func (q *LinkedListQueue) Dequeue() interface{} {
	if q.front == nil {
		return nil
	}
	item := q.front.Val
	q.front = q.front.Next
	if q.front == nil {
		q.rear = nil
	}
	return item
}

func (q *LinkedListQueue) Size() int {
	size := 0
	node := q.front
	for node != nil {
		size++
		node = node.Next
	}
	return size
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.front == nil
}
