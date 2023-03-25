package main

import (
	"errors"
	"fmt"
)

type node struct {
	value interface{}
	prev  *node
	next  *node
}

type deque struct {
	length int
	head   *node
	tail   *node
}

func (d *deque) isEmpty() bool {
	return d.length == 0
}

func (d *deque) len() int {
	return d.length
}

func (d *deque) pushFront(val interface{}) {
	n := &node{val, nil, d.head}
	if d.head != nil {
		d.head.prev = n
	}
	d.head = n
	if d.tail == nil {
		d.tail = n
	}
	d.length++
}

func (d *deque) pushBack(val interface{}) {
	n := &node{val, d.tail, nil}
	if d.tail != nil {
		d.tail.next = n
	}
	d.tail = n
	if d.head == nil {
		d.head = n
	}
	d.length++
}

func (d *deque) popFront() (interface{}, error) {
	if d.isEmpty() {
		return nil, errors.New("deque is empty")
	}
	n := d.head
	d.head = n.next
	if d.head != nil {
		d.head.prev = nil
	} else {
		d.tail = nil
	}
	d.length--
	return n.value, nil
}

func (d *deque) popBack() (interface{}, error) {
	if d.isEmpty() {
		return nil, errors.New("deque is empty")
	}
	n := d.tail
	d.tail = n.prev
	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.head = nil
	}
	d.length--
	return n.value, nil
}

func main() {
	d := &deque{}
	d.pushBack(1)
	d.pushFront(2)
	d.pushBack(3)
	d.pushFront(4)
	fmt.Println(d.len()) // 4
	for !d.isEmpty() {
		val, _ := d.popFront()
		fmt.Println(val)
	}
}

// 4 2 1 3
