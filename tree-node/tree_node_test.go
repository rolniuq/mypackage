package treenode

import (
	"testing"
)

func Test_Create(t *testing.T) {
	values := []int{3, 9, 20, -1, -1, 15, 7}

	tree := &TreeNode[int]{}
	tree = tree.Create(values)

	tree.Print()
}
