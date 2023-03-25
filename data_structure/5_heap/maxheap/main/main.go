package main

import "fmt"

type MaxHeap struct {
	data []int // 用数组存储堆元素
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{[]int{}}
}

// 返回父节点的索引
func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

// 返回左子节点的索引
func (h *MaxHeap) leftChild(i int) int {
	return 2*i + 1
}

// 返回右子节点的索引
func (h *MaxHeap) rightChild(i int) int {
	return 2*i + 2
}

// 向堆中插入一个元素
func (h *MaxHeap) Insert(key int) {
	h.data = append(h.data, key) // 将元素添加到数组末尾
	i := len(h.data) - 1         // 新元素的索引
	for i > 0 && h.data[i] > h.data[h.parent(i)] {
		// 如果新元素比父节点大，则交换它们的位置
		h.data[i], h.data[h.parent(i)] = h.data[h.parent(i)], h.data[i]
		i = h.parent(i) // 更新索引
	}
}

// 从堆中删除最大元素
func (h *MaxHeap) ExtractMax() int {
	max := h.data[0]                  // 最大元素为根节点
	h.data[0] = h.data[len(h.data)-1] // 将最后一个元素移到根节点
	h.data = h.data[:len(h.data)-1]   // 删除最后一个元素
	h.maxHeapify(0)                   // 从根节点开始进行堆化操作
	return max
}

// 将指定的节点进行堆化操作
func (h *MaxHeap) maxHeapify(i int) {
	left := h.leftChild(i)
	right := h.rightChild(i)
	largest := i
	if left < len(h.data) && h.data[left] > h.data[largest] {
		largest = left
	}
	if right < len(h.data) && h.data[right] > h.data[largest] {
		largest = right
	}
	if largest != i {
		// 如果当前节点不是最大的，则交换它和最大的子节点的位置
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		h.maxHeapify(largest) // 递归进行堆化操作
	}
}

func main() {
	h := NewMaxHeap()
	h.Insert(5)
	h.Insert(2)
	h.Insert(7)
	h.Insert(3)
	fmt.Println(h.data)         // [7 3 2 5]
	fmt.Println(h.ExtractMax()) // 7
	fmt.Println(h.data)         // [5 3 2]
}
