package search

import (
	"fmt"
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("before binary search of a:", a, "ready to find element is:", 7)
	index := BinarySearch(a, 7, 0, 9)
	fmt.Println("after binary search of a,index is:", index)
}
