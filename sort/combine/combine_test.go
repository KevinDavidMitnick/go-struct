package combine

import (
	"fmt"
	"testing"
)

func Test_CombineInsert(t *testing.T) {
	a := []int{1, 3, 9, 4, 2, 6, 5, 7, 8, 0}
	fmt.Println("before combine of a:", a)
	b := Combine(a)
	fmt.Println("after combine of a:", b)
}
