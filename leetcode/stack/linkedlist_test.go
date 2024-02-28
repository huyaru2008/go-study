package stack

import "testing"

func TestListPush(t *testing.T) {
	stack := NewStackList()
	stack.Push("q")
	stack.Push("w")
	stack.Push("e")
	t.Logf("%s", stack.Pop())
}
