package main

import (
	"container/heap"
	"fmt"
	"time"
)

type Item struct {
	value    string
	priority int
	index    int
	expire   time.Time
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value string, priority int, expire time.Time) {
	item.value = value
	item.priority = priority
	item.expire = expire
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) cleanExpired() {
	for pq.Len() > 0 {
		item := (*pq)[0]
		if item.expire.After(time.Now()) {
			return
		}
		heap.Pop(pq)
	}
}

func main() {
	pq := make(PriorityQueue, 0)

	// Insert an item with an expiration time of 10 seconds
	item := &Item{
		value:    "foo",
		priority: 1,
		expire:   time.Now().Add(10 * time.Second),
	}
	heap.Push(&pq, item)

	// Update the item with a new value, priority and expiration time
	pq.update(item, "bar", 2, time.Now().Add(20*time.Second))

	// Clean up any expired items
	pq.cleanExpired()

	// Pop the item with the highest priority
	item = heap.Pop(&pq).(*Item)
	fmt.Printf("value=%s priority=%d\n", item.value, item.priority)
}
