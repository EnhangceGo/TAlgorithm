package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (list *LinkedList) Append(data interface{}) {
	node := &Node{data, nil}
	if list.head == nil {
		list.head = node
	} else {
		cur := list.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node
	}
	list.length++
}

func (list *LinkedList) Insert(index int, data interface{}) error {
	if index < 0 || index > list.length {
		return fmt.Errorf("Index out of range")
	}
	node := &Node{data, nil}
	if index == 0 {
		node.next = list.head
		list.head = node
	} else {
		cur := list.head
		for i := 0; i < index-1; i++ {
			cur = cur.next
		}
		node.next = cur.next
		cur.next = node
	}
	list.length++
	return nil
}

func (list *LinkedList) Delete(index int) error {
	if index < 0 || index >= list.length {
		return fmt.Errorf("Index out of range")
	}
	if index == 0 {
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

func (list *LinkedList) Search(data interface{}) (int, error) {
	cur := list.head
	for i := 0; cur != nil; i++ {
		if cur.data == data {
			return i, nil
		}
		cur = cur.next
	}
	return -1, fmt.Errorf("Data not found")
}

func (list *LinkedList) Print() {
	cur := list.head
	for cur != nil {
		fmt.Printf("%v ", cur.data)
		cur = cur.next
	}
	fmt.Println()
}

func main() {
	list := &LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Print() // 1 2 3 4

	list.Insert(0, 0)
	list.Insert(5, 5)
	list.Print() // 0 1 2 3 4 5

	list.Delete(0)
	list.Delete(4)
	list.Print() // 1 2 3 4

	index, err := list.Search(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data found at index", index)
	}
}
