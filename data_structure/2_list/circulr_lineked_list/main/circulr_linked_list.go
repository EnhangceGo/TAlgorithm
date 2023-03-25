package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type CircularLinkedList struct {
	head   *Node
	length int
}

func (list *CircularLinkedList) Append(data interface{}) {
	node := &Node{data, nil}
	if list.head == nil {
		node.next = node
		list.head = node
	} else {
		cur := list.head
		for cur.next != list.head {
			cur = cur.next
		}
		cur.next = node
		node.next = list.head
	}
	list.length++
}

func (list *CircularLinkedList) Prepend(data interface{}) {
	node := &Node{data, nil}
	if list.head == nil {
		node.next = node
	} else {
		cur := list.head
		for cur.next != list.head {
			cur = cur.next
		}
		cur.next = node
		node.next = list.head
		list.head = node
	}
	list.length++
}

func (list *CircularLinkedList) Remove(index int) error {
	if index < 0 || index >= list.length {
		return fmt.Errorf("Index out of range")
	}
	if index == 0 {
		cur := list.head
		for cur.next != list.head {
			cur = cur.next
		}
		cur.next = list.head.next
		list.head = list.head.next
	} else {
		cur := list.head
		for i := 0; i < index-1; i++ {
			cur = cur.next
		}
		cur.next = cur.next.next
	}
	list.length--
	return nil
}

func (list *CircularLinkedList) Traverse() []interface{} {
	var result []interface{}
	if list.head == nil {
		return result
	}
	cur := list.head
	for {
		result = append(result, cur.data)
		cur = cur.next
		if cur == list.head {
			break
		}
	}
	return result
}

func main() {
	list := CircularLinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Prepend(0)
	fmt.Println(list.Traverse())
	list.Remove(2)
	fmt.Println(list.Traverse())
}
