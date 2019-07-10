package huffmantree

import (
	"fmt"
)

// BinaryTree , when parent is nil,it is empty
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

func (tree *BinaryTree) IsEmpty() bool {
	if tree.parent == nil {
		return true
	}
	return false
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

// FindNode if we find return data,else return -1
func (tree *BinaryTree) FindNode(data int) *BinaryTree {
	if tree == nil {
		return nil
	}
	if data == tree.data {
		return tree
	}

	if data > tree.data {
		return tree.right.FindNode(data)
	}

	if data < tree.data {
		return tree.left.FindNode(data)
	}
	return nil
}

// RemoveData remove node data math
func (tree *BinaryTree) RemoveData(data int) {
	found := tree.FindNode(data)
	if found == nil {
		return
	}
	if found.left == nil && found.right == nil {
		tree.RemoveLeafNode(found)
	} else if found.left == nil {
		found.data = found.right.data
		found.left = found.right.left
		found.right = found.right.right
		if found.left != nil {
			found.left.parent = found
		}
		if found.right != nil {
			found.right.parent = found
		}
	} else if found.right == nil {
		found.data = found.left.data
		found.right = found.left.right
		found.left = found.left.left
		if found.left != nil {
			found.left.parent = found
		}
		if found.right != nil {
			found.right.parent = found
		}
	} else {
		max := found.left.FindMaxNode()
		if max != nil {
			found.data = max.data
			tree.RemoveLeafNode(max)
		}
	}

}

// FindMaxNode find max node
func (tree *BinaryTree) FindMaxNode() *BinaryTree {
	if tree.left == nil && tree.right == nil {
		return tree
	}
	if tree.right == nil {
		return tree
	}
	return tree.right.FindMaxNode()

}

// RemoveLeafNode remove leaf node
func (tree *BinaryTree) RemoveLeafNode(node *BinaryTree) {
	// remove the root node
	if node.parent == node {
		tree.parent = nil
		return
	}
	parent := node.parent
	if parent.right == node {
		parent.right = nil
	}
	if parent.left == node {
		parent.left = nil
	}

}

func (tree *BinaryTree) PreOrder(callback func(int)) {
	if tree == nil || tree.parent == nil {
		return
	}
	callback(tree.data)
	tree.left.PreOrder(callback)
	tree.right.PreOrder(callback)
}

func (tree *BinaryTree) MidOrder(callback func(int)) {
	if tree == nil || tree.parent == nil {
		return
	}
	tree.left.MidOrder(callback)
	callback(tree.data)
	tree.right.MidOrder(callback)
}

func (tree *BinaryTree) PostOrder(callback func(int)) {
	if tree == nil || tree.parent == nil {
		return
	}
	tree.left.PostOrder(callback)
	tree.right.PostOrder(callback)
	callback(tree.data)

}
