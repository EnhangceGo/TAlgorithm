package main

import (
	"fmt"
	"math/rand"
)

const (
	maxLevel = 16  // SkipList 的最大层数
	p        = 0.5 // 节点向上提升一层的概率
)

type node struct {
	key     int         // 节点的 key
	value   interface{} // 节点的 value
	forward []*node     // 节点每一层的后继指针
}

type SkipList struct {
	head   *node // SkipList 的头节点
	level  int   // SkipList 的层数
	length int   // SkipList 的长度
}

// newNode 创建一个新节点
func newNode(key int, value interface{}, level int) *node {
	return &node{
		key:     key,
		value:   value,
		forward: make([]*node, level),
	}
}

// NewSkipList 创建一个新的 SkipList
func NewSkipList() *SkipList {
	return &SkipList{
		head:   newNode(0, nil, maxLevel),
		level:  1,
		length: 0,
	}
}

// randomLevel 随机生成节点的层数
func randomLevel() int {
	level := 1
	for rand.Float64() < p && level < maxLevel {
		level++
	}
	return level
}

// Insert 将一个节点插入到 SkipList 中
func (sl *SkipList) Insert(key int, value interface{}) {
	update := make([]*node, maxLevel)
	x := sl.head
	// 从高到低遍历每一层，找到要插入的位置
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
		update[i] = x
	}
	// 如果 key 已经存在，更新对应的 value
	x = x.forward[0]
	if x != nil && x.key == key {
		x.value = value
	} else {
		// 随机生成节点的层数
		level := randomLevel()
		if level > sl.level {
			// 如果节点的层数大于 SkipList 的层数，更新 SkipList 的层数
			for i := sl.level; i < level; i++ {
				update[i] = sl.head
			}
			sl.level = level
		}
		// 创建新节点
		x = newNode(key, value, level)
		// 更新每一层的后继指针
		for i := 0; i < level; i++ {
			x.forward[i] = update[i].forward[i]
			update[i].forward[i] = x
		}
		sl.length++
	}
}

// Delete 从 SkipList 中删除一个节点
func (sl *SkipList) Delete(key int) {
	update := make([]*node, maxLevel)
	x := sl.head
	// 从高到低遍历每一层，找到要删除的节点
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
		// 更新 update 数组，记录每一层需要更新的节点
		update[i] = x
	}
	x = x.forward[0]
	if x != nil && x.key == key {
		// 如果找到了要删除的节点，更新每一层的后继指针
		for i := 0; i < sl.level; i++ {
			if update[i].forward[i] != x {
				break
			}
			update[i].forward[i] = x.forward[i]
		}
		sl.length--
		// 如果删除了最高层的节点，更新 SkipList 的层数
		for sl.level > 1 && sl.head.forward[sl.level-1] == nil {
			sl.level--
		}
	}
}

// Search 在 SkipList 中查找一个节点
func (sl *SkipList) Search(key int) interface{} {
	x := sl.head
	// 从高到低遍历每一层，找到要查找的节点
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.key == key {
		// 如果找到了要查找的节点，返回节点的 value
		return x.value
	}
	// 没有找到要查找的节点，返回 nil
	return nil
}
func main() {
	sl := NewSkipList()
	sl.Insert(3, "value1")
	sl.Insert(1, "value2")
	sl.Insert(2, "value3")
	fmt.Println(sl.Search(1)) // value2
	fmt.Println(sl.Search(2)) // value3
	fmt.Println(sl.Search(3)) // value1

	sl.Delete(2)

	fmt.Println(sl.Search(2)) // nil
}
