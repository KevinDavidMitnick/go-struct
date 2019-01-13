package swap

import (
	"fmt"
	"testing"
)

func Test_FastSort(t *testing.T) {
	a := []int{1, 3, 9, 4, 2, 6, 5, 7, 8, 0}
	fmt.Println("before fast sort of a:", a)
	b := FastSort(a)
	fmt.Println("after fast sort of a:", b)
}
