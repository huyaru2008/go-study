package stack

import "testing"

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	t.Logf("%s", stack.Pop())
}
