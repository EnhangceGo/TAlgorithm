package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

// 双向链表节点
type node struct {
	val  int
	prev unsafe.Pointer // 前驱指针，使用 unsafe.Pointer 类型，保证指针操作的原子性
	next unsafe.Pointer // 后继指针，使用 unsafe.Pointer 类型，保证指针操作的原子性
}

// 双向链表
type list struct {
	head unsafe.Pointer // 头节点指针，使用 unsafe.Pointer 类型，保证指针操作的原子性
	tail unsafe.Pointer // 尾节点指针，使用 unsafe.Pointer 类型，保证指针操作的原子性
	len  int32          // 链表长度，使用 int32 类型，保证原子操作
}

// 初始化双向链表
func newList() *list {
	n := unsafe.Pointer(new(node))
	return &list{
		head: n,
		tail: n,
		len:  0,
	}
}

// 头部添加节点
func (l *list) pushFront(val int) {
	newHead := unsafe.Pointer(&node{
		val:  val,
		prev: nil,
		next: l.head,
	})
	oldHead := atomic.SwapPointer(&l.head, newHead)
	(*node)(oldHead).prev = newHead
	atomic.AddInt32(&l.len, 1)
}

// 尾部添加节点
func (l *list) pushBack(val int) {
	newTail := unsafe.Pointer(&node{
		val:  val,
		prev: l.tail,
		next: nil,
	})
	for {
		oldTail := atomic.LoadPointer(&l.tail)
		if atomic.CompareAndSwapPointer(&(*node)(oldTail).next, nil, newTail) {
			if atomic.CompareAndSwapPointer(&l.tail, oldTail, newTail) {
				(*node)(oldTail).next = newTail
				atomic.AddInt32(&l.len, 1)
				return
			}
		}
	}
}

// 头部弹出节点
func (l *list) popFront() int {
	for {
		oldHead := atomic.LoadPointer(&l.head)
		if oldHead == atomic.LoadPointer(&l.tail) {
			return -1
		}
		newHead := (*node)(oldHead).next
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&l.head)), oldHead, newHead) {
			(*node)(newHead).prev = nil
			val := (*node)(oldHead).val
			atomic.AddInt32(&l.len, -1)
			return val
		}
	}
}

// 尾部弹出节点
func (l *list) popBack() int {
	for {
		oldTail := atomic.LoadPointer(&l.tail)
		newTail := (*node)(oldTail).prev
		if oldTail == atomic.LoadPointer(&l.head) {
			return -1
		}
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&l.tail)), oldTail, newTail) {
			(*node)(newTail).next = nil
			val := (*node)(oldTail).val
			atomic.AddInt32(&l.len, -1)
			return val
		}
	}
}

// 获取链表长度
func (l *list) lenList() int {
	return int(atomic.LoadInt32(&l.len))
}

func main() {
	l := newList()
	l.pushFront(1)
	l.pushFront(2)
	l.pushFront(3)
	l.pushBack(4)
	l.pushBack(5)
	l.pushBack(6)
	fmt.Println(l.popFront())
	fmt.Println(l.popFront())
	fmt.Println(l.popBack())
	fmt.Println(l.popBack())
	fmt.Println(l.lenList())
}
