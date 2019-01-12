package insert

import (
	"fmt"
	"testing"
)

func Test_SimpleInsert(t *testing.T) {
	a := []int{1, 3, 9, 4, 2, 6, 5, 7, 8, 0}
	fmt.Println("before simple insert of a:", a)
	b := SimpleInsert(a)
	fmt.Println("after simple insert of a:", b)
}
