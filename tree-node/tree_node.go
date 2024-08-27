package treenode

import "fmt"

type TreeNode[T int] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func (t *TreeNode[T]) Create(values []T) *TreeNode[T] {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode[T]{Val: values[0]}
	queue := []*TreeNode[T]{root}

	i := 1
	for i < len(values) {
		current := queue[0]
		queue = queue[1:]

		// Create left child
		if i < len(values) {
			current.Left = &TreeNode[T]{Val: values[i]}
			queue = append(queue, current.Left)
			i++
		}

		// Create right child
		if i < len(values) {
			current.Right = &TreeNode[T]{Val: values[i]}
			queue = append(queue, current.Right)
			i++
		}
	}

	return root
}

func (t *TreeNode[T]) Print() {
	if t == nil {
		return
	}

	t.Left.Print()

	fmt.Print(t.Val, " ")

	t.Right.Print()
}
