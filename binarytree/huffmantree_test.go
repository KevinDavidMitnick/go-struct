package huffmantree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateHuffmanTree(t *testing.T) {
	tree := CreateHuffmanTree()
	assert.NotEqual(t, nil, tree)
}

func Test_InsertNode(t *testing.T) {
	tree := CreateHuffmanTree()
	tree.InsertNode(6)
	tree.InsertNode(3)
	tree.InsertNode(2)
	tree.InsertNode(7)
	tree.InsertNode(5)
	tree.InsertNode(8)

	fmt.Print("preorder:")
	tree.PreOrder(Println)
	fmt.Println()

	fmt.Print("midorder:")
	tree.MidOrder(Println)
	fmt.Println()

	fmt.Print("postorder:")
	tree.PostOrder(Println)
	fmt.Println()

	assert.NotEqual(t, nil, tree)
}
