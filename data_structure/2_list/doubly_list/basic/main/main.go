package main

import "fmt"

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (list *DoublyLinkedList) Append(data interface{}) {
	node := &Node{data, nil, nil}
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.prev = list.tail
		list.tail.next = node
		list.tail = node
	}
	list.length++
}

func (list *DoublyLinkedList) Prepend(data interface{}) {
	node := &Node{data, nil, nil}
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head.prev = node
		list.head = node
	}
	list.length++
}

func (list *DoublyLinkedList) Remove(index int) error {
	if index < 0 || index >= list.length {
		return fmt.Errorf("Index out of range")
	}
	if index == 0 {
		list.head = list.head.next
		if list.head == nil {
			list.tail = nil
		} else {
			list.head.prev = nil
		}
	} else if index == list.length-1 {
		list.tail = list.tail.prev
		list.tail.next = nil
	} else {
		cur := list.head
		for i := 0; i < index; i++ {
			cur = cur.next
		}
		cur.prev.next = cur.next
		cur.next.prev = cur.prev
	}
	list.length--
	return nil
}

func (list *DoublyLinkedList) TraverseFromHead() []interface{} {
	var result []interface{}
	cur := list.head
	for cur != nil {
		result = append(result, cur.data)
		cur = cur.next
	}
	return result
}

func (list *DoublyLinkedList) TraverseFromTail() []interface{} {
	var result []interface{}
	cur := list.tail
	for cur != nil {
		result = append(result, cur.data)
		cur = cur.prev
	}
	return result
}

func main() {
	list := DoublyLinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Prepend(0)
	fmt.Println(list.TraverseFromHead())
	fmt.Println(list.TraverseFromTail())
	list.Remove(2)
	fmt.Println(list.TraverseFromHead())
	fmt.Println(list.TraverseFromTail())
}
