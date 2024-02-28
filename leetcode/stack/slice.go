package stack

/*
数组实现顺序栈
*/
type StackSlice struct {
	items []string
	index int
}

func NewStack() *StackSlice {
	return &StackSlice{
		items: make([]string, 10),
		index: 0,
	}
}

// Push 入栈
func (s *StackSlice) Push(item string) {
	s.index++
	if s.index == len(s.items) {
		s.items = append(s.items, item)
		return
	}
	s.items[s.index] = item
}

// Pop 出栈
func (s *StackSlice) Pop() string {
	if s.index == 0 {
		return ""
	}
	item := s.items[s.index]
	s.index--
	return item
}
