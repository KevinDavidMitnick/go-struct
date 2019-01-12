package swap

import (
	"fmt"
	"testing"
)

func Test_BubbleBigSort(t *testing.T) {
	a := []int{1, 3, 9, 4, 2, 6, 5, 7, 8, 0}
	fmt.Println("before big sort of a:", a)
	b := BubbleBigSort(a)
	fmt.Println("after big sort of a:", b)
}

func Test_BubbleSmallSort(t *testing.T) {
	a := []int{1, 3, 9, 4, 2, 6, 5, 7, 8, 0}
	fmt.Println("before small sort of a:", a)
	b := BubbleSmallSort(a)
	fmt.Println("after small sort of a:", b)
}
