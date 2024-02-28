package stack

/*
链表链式栈:双向循环链表

*/

type node struct {
	prev, next *node
	value      string
}
type StackList struct {
	root *node
	len  int
}

func NewStackList() *StackList {
	n := &node{}
	n.next = n
	n.prev = n
	return &StackList{
		root: n,
		len:  0,
	}
}

func (s *StackList) Push(value string) {
	n := &node{value: value}
	s.root.prev.next = n
	n.prev = s.root.prev
	n.next = s.root
	s.root.prev = n
	s.len++
}

func (s *StackList) Pop() string {
	item := s.root.prev
	if item == s.root {
		return ""
	}

	s.root.prev = item.prev
	item.prev.next = s.root
	item.prev = nil
	item.next = nil
	s.len--
	return item.value
}
