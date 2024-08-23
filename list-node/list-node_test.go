package listnode

import "testing"

func TestListNode_Print(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	l := &ListNode{}
	l = l.Create(values)
	l.Print()
}
