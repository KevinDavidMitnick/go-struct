package insert

import (
	"fmt"
	"testing"
)

func Test_ShellInsert(t *testing.T) {
	a := []int{1, 3, 9, 4, 2, 6, 5, 7, 8, 0}
	fmt.Println("before shell insert of a:", a)
	b := ShellInsert(a)
	fmt.Println("after shell insert of a:", b)
}
