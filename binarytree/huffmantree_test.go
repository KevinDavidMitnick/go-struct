package huffmantree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func Test_FindNode(t *testing.T) {
	tree := CreateHuffmanTree()
	tree.InsertNode(6)
	tree.InsertNode(3)
	tree.InsertNode(2)
	tree.InsertNode(7)
	tree.InsertNode(7)
	tree.InsertNode(5)
	tree.InsertNode(8)

	found := tree.FindNode(4)
	fmt.Println(found)

	found = tree.FindNode(7)
	fmt.Println(found.data)
	assert.Equal(t, found.data, 7)

	found = tree.FindNode(8)
	fmt.Println(found.data)
	assert.Equal(t, found.data, 8)

	found = tree.FindNode(6)
	fmt.Println(found.data)
	assert.Equal(t, found.data, 6)

	max := tree.FindMaxNode()
	fmt.Println("find max node :", max.data)
	assert.Equal(t, max.data, 8)

	tree.RemoveData(7)
	fmt.Print("midorder:")
	tree.MidOrder(Println)
	fmt.Println()

	tree.RemoveData(2)
	fmt.Print("midorder:")
	tree.MidOrder(Println)
	fmt.Println()

	tree.RemoveData(8)
	fmt.Print("midorder:")
	tree.MidOrder(Println)
	fmt.Println()

	tree.RemoveData(5)
	fmt.Print("midorder:")
	tree.MidOrder(Println)
	fmt.Println()

	tree.RemoveData(6)
	fmt.Print("midorder:")
	tree.MidOrder(Println)
	fmt.Println()

	tree.RemoveData(7)
	fmt.Print("midorder:")
	tree.MidOrder(Println)
	fmt.Println()

	tree.RemoveData(3)
	fmt.Print("midorder:")
	tree.MidOrder(Println)
	fmt.Println()

	tree.RemoveData(8)
	fmt.Println("after remove.")
	fmt.Print("midorder:")
	tree.MidOrder(Println)
	fmt.Println()
}
