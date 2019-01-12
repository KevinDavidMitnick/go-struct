package insert

import (
	"fmt"
	"testing"
)

func Test_BinaryInsert(t *testing.T) {
	a := []int{1, 3, 9, 4, 2, 6, 5, 7, 8, 0}
	fmt.Println("before binary insert of a:", a)
	b := BinaryInsert(a)
	fmt.Println("after binary insert of a:", b)
}
