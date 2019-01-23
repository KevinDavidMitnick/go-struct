package search

import (
	"fmt"
	"testing"
)

func Test_NormalSearch(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("before normal search of a:", a, "ready to find element is:", 7)
	index := NormalSearch(a, 7)
	fmt.Println("after normal search of a,index is:", index)
}
