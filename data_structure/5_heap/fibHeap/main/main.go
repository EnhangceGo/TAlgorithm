package main

import (
	"container/heap"
	"fmt"
)

// 存储节点数据的结构体
type Node struct {
	value    int     // 节点的值
	children []*Node // 子节点列表
	parent   *Node   // 父节点
}

// 返回一个空的斐波那契堆
func NewHeap() *Heap {
	return new(Heap)
}

// 堆类型，包含一个表示节点的切片，并重写 container/heap 的接口
type Heap struct {
	nodes []*Node
}

// 获取堆中节点的数量
func (h *Heap) Len() int {
	return len(h.nodes)
}

// 比较 i、j 两个节点的大小，用于决定由 container/heap 接口调用的插入和删除顺序
func (h *Heap) Less(i, j int) bool {
	return h.nodes[i].value < h.nodes[j].value
}

// 交换节点
func (h *Heap) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

// 向堆中插入一个节点
func (h *Heap) Push(x interface{}) {
	h.nodes = append(h.nodes, x.(*Node))
}

// 从堆中删除最小的节点，并将其返回
func (h *Heap) Pop() interface{} {
	oldNodes := h.nodes
	n := len(oldNodes)
	node := oldNodes[n-1] // 获取最后一个节点
	h.nodes = oldNodes[0 : n-1]
	return node
}

// 合并两个斐波那契堆
func (h *Heap) Merge(h2 *Heap) {
	heap.Push(h, heap.Pop(h2))
	for _, node := range h2.nodes {
		heap.Push(h, node)
	}
}

// 向斐波那契堆中插入一个节点
func (h *Heap) Insert(value int) *Node {
	node := &Node{value: value}
	heap.Push(h, node)
	return node
}

// 提取堆中最小的节点
func (h *Heap) ExtractMin() *Node {
	node := heap.Pop(h).(*Node)
	for _, child := range node.children {
		child.parent = nil // 将 "父节点" 置为空，以表示它成为一个独立的节点
		heap.Push(h, child)
	}
	return node
}

// 将一个节点加入到斐波那契堆中
func (h *Heap) InsertNode(node *Node) {
	heap.Push(h, node)
}

func main() {
	h := NewHeap()

	// 向堆中插入 8 个节点
	h.Insert(2)
	h.Insert(3)
	h.Insert(1)
	h.Insert(7)
	h.Insert(8)
	h.Insert(6)
	h.Insert(0)
	h.Insert(5)

	// 输出堆中最小的节点
	fmt.Println(h.ExtractMin().value)

	// 输出堆中次小的节点
	fmt.Println(h.ExtractMin().value)

	// 提取两次最小的节点之后，向堆中插入一些节点
	h.InsertNode(&Node{value: 4})
	h.InsertNode(&Node{value: 1})
	h.InsertNode(&Node{value: 2})

	// 输出堆中最小的节点
	fmt.Println(h.ExtractMin().value)
}
