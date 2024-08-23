package listnode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Create(values []int) *ListNode {
	if l == nil {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		newNode := &ListNode{Val: values[i]}
		current.Next = newNode
		current = newNode
	}

	return head
}

func (l *ListNode) Print() {
	if l == nil {
		return
	}

	if l.Next != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
