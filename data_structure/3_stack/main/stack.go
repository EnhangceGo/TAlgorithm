package main

import "fmt"

type Stack struct {
	top  int
	data []int
}

func NewStack() *Stack {
	return &Stack{
		top:  -1,
		data: make([]int, 0),
	}
}

func (s *Stack) Push(val int) {
	s.top++
	if len(s.data) > s.top {
		s.data[s.top] = val
	} else {
		s.data = append(s.data, val)
	}
}

func (s *Stack) Pop() int {
	if s.top == -1 {
		panic("stack is empty")
	}
	val := s.data[s.top]
	s.top--
	return val
}

func (s *Stack) Peek() int {
	if s.top == -1 {
		panic("stack is empty")
	}
	return s.data[s.top]
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Size() int {
	return s.top + 1
}

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("栈顶元素：", stack.Peek())
	fmt.Println("出栈元素：", stack.Pop())
	fmt.Println("出栈元素：", stack.Pop())
	fmt.Println("栈顶元素：", stack.Peek())
	fmt.Println("栈是否为空：", stack.IsEmpty())
	fmt.Println("栈中元素个数：", stack.Size())
}
