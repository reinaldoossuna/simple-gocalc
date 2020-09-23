package stack

import (
	"fmt"
)

type Node struct {
	Value float64
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

type Stack struct {
	head *Node // first off
	count int
}
func (s *Stack) Count() int {
	return s.count
}

func (s *Stack) Push(v float64) {
	n := Node{v, s.head}
	s.count++
	s.head = &n
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--

	head := s.head
	s.head = head.next
	return head
}

func (s *Stack) Shuffle() {
	if s.count < 1 {
		return
	}

	oldHead := s.head
	newHead := oldHead.next

	oldHead.next = newHead.next
	newHead.next = oldHead
	s.head = newHead
}

func (s *Stack) Pick(n int) *Node {
	if n > s.count {
		return nil
	}

	oldHead := s.head
	temp := s.head.next
	for i := 0; i < n - 1; i++ {
		temp = temp.next
	}


	newHead := temp
	newHead.next = oldHead
	s.head = newHead
	return s.Pop()
}

func (s *Stack) Preview(n int) []float64 {
	var l []float64
	temp := s.head
	for i := 0; i < s.count && i < n; i++ {
		l = append(l, temp.Value)
		temp = temp.next
	}
	return l
}
