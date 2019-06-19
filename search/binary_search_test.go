package search

import (
	"fmt"
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
	fmt.Println("before binary search of a:", a, "ready to find element is:", 7)
	index := Search1(a, 7, 0, len(a)-1)
	fmt.Println("after binary search of a,index is:", index)

	fmt.Println("before binary search of a:", a, "ready to find element is:", 7)
	index = Search2(a, 8, 0, len(a)-1)
	fmt.Println("after binary search of a,index is:", index)

	fmt.Println("before binary search of a:", a, "ready to find element is:", 7)
	index = Search2(a, 5, 0, len(a)-1)
	fmt.Println("after binary search of a,index is:", index)
}
