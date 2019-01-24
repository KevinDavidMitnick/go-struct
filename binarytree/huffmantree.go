package huffmantree

import (
	"fmt"
)

type BinaryTree struct {
	right  *BinaryTree
	data   int
	parent *BinaryTree
	left   *BinaryTree
}

func Println(data int) {
	fmt.Print(data, " ")
}

func CreateHuffmanTree() *BinaryTree {
	return &BinaryTree{}
}

func (tree *BinaryTree) InsertNode(data int) {
	if tree.parent == nil {
		tree.data = data
		tree.parent = tree
		return
	}
	if data >= tree.data {
		if tree.right == nil {
			tree.right = &BinaryTree{data: data, parent: tree}
		} else {
			tree.right.InsertNode(data)
		}
	}
	if data < tree.data {
		if tree.left == nil {
			tree.left = &BinaryTree{data: data, parent: tree}
		} else {
			tree.left.InsertNode(data)
		}
	}
}

func (tree *BinaryTree) PreOrder(callback func(int)) {
	if tree == nil {
		return
	}
	callback(tree.data)
	tree.left.PreOrder(callback)
	tree.right.PreOrder(callback)
}

func (tree *BinaryTree) MidOrder(callback func(int)) {
	if tree == nil {
		return
	}
	tree.left.MidOrder(callback)
	callback(tree.data)
	tree.right.MidOrder(callback)
}

func (tree *BinaryTree) PostOrder(callback func(int)) {
	if tree == nil {
		return
	}
	tree.left.PostOrder(callback)
	tree.right.PostOrder(callback)
	callback(tree.data)

}
