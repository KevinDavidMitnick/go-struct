package selector

import (
	"fmt"
	"testing"
)

func Test_SimpleSelect(t *testing.T) {
	a := []int{1, 3, 9, 4, 2, 6, 5, 7, 8, 0}
	fmt.Println("before simple select of a:", a)
	b := SimpleSelect(a)
	fmt.Println("after simple sort of a:", b)
}
