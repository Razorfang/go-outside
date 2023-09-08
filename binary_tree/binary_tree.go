package special_binary_tree

import (
	"fmt"
)

type Tree[T any] struct {
	root *TreeNode[T]
	nodeCount uint64
}

type TreeNode[T any] struct {
	value T
	leftNode *TreeNode[T]
	rightNode *TreeNode[T]
}

/* In order for an element to exist in a binary tree, we must be able to compare it to another element */
type TreeElement[T any] interface {
	compare(a TreeNode[T], b TreeNode[T]) int
}

func newTree[T any]() *Tree[T] {
	return &Tree[T]{nil, 0}
}

func test() {
	tree := newTree[float64]()
	fmt.Println(tree.nodeCount, tree.root)
}
