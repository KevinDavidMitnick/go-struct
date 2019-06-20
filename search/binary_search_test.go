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

	fmt.Println("before binary search of a:", a, "ready to find element is:", 5)
	index = Search2(a, 5, 0, len(a)-1)
	fmt.Println("after binary search of a,index is:", index)

	a = []int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4}
	fmt.Println("before CircleSearch search of a:", a, "ready to find element is:", 0)
	index = CircleSearch(a, 0, 0, len(a)-1)

	fmt.Println("after CircleSearch search of a,index is:", index)
	fmt.Println("before CircleSearch search of a:", a, "ready to find element is:", 1)
	index = CircleSearch(a, 1, 0, len(a)-1)

	fmt.Println("after CircleSearch search of a,index is:", index)
	fmt.Println("before CircleSearch search of a:", a, "ready to find element is:", 7)
	index = CircleSearch(a, 7, 0, len(a)-1)

	fmt.Println("after CircleSearch search of a,index is:", index)
	fmt.Println("before CircleSearch search of a:", a, "ready to find element is:", 6)
	index = CircleSearch(a, 6, 0, len(a)-1)

	fmt.Println("after CircleSearch search of a,index is:", index)
	fmt.Println("before CircleSearch search of a:", a, "ready to find element is:", 8)
	index = CircleSearch(a, 8, 0, len(a)-1)

	fmt.Println("after CircleSearch search of a,index is:", index)
	fmt.Println("before CircleSearch search of a:", a, "ready to find element is:", 9)
	index = CircleSearch(a, 9, 0, len(a)-1)

	fmt.Println("after CircleSearch search of a,index is:", index)
	fmt.Println("before CircleSearch search of a:", a, "ready to find element is:", 3)
	index = CircleSearch(a, 3, 0, len(a)-1)

	fmt.Println("after CircleSearch search of a,index is:", index)
	fmt.Println("before CircleSearch search of a:", a, "ready to find element is:", 4)
	index = CircleSearch(a, 4, 0, len(a)-1)

	fmt.Println("after CircleSearch search of a,index is:", index)
}
