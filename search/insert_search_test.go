package search

import (
	"fmt"
	"testing"
)

func Test_InsertSearch(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("before insert search of a:", a, "ready to find element is:", 7)
	index := InsertSearch(a, 7, 0, 9)
	fmt.Println("after insert search of a,index is:", index)
}
