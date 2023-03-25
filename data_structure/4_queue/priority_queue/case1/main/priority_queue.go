package main

import (
	"container/heap"
	"fmt"
)

type priorityQueue struct {
	heap []int
}

func (p *priorityQueue) Len() int { return len(p.heap) }

func (p *priorityQueue) Less(i, j int) bool { return p.heap[i] > p.heap[j] }

func (p *priorityQueue) Swap(i, j int) { p.heap[i], p.heap[j] = p.heap[j], p.heap[i] }

func (p *priorityQueue) Push(x interface{}) {
	p.heap = append(p.heap, x.(int))
}

func (p *priorityQueue) Pop() interface{} {
	n := len(p.heap)
	x := p.heap[n-1]
	p.heap = p.heap[:n-1]
	return x
}

func main() {
	pq := &priorityQueue{}
	heap.Init(pq)

	heap.Push(pq, 3)
	heap.Push(pq, 1)
	heap.Push(pq, 4)
	heap.Push(pq, 1)

	for pq.Len() > 0 {
		fmt.Println(heap.Pop(pq))
	}
	// 4 3 1 1
}
