package stack

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InitArrayStack(t *testing.T) {
	a := -1
	st, err := InitArrayStack(a)
	assert.Equal(t, nil, st)
	assert.Equal(t, errors.New("list size less than 1"), err)

	a = 0
	st, err = InitArrayStack(a)
	assert.Equal(t, nil, st)
	assert.Equal(t, errors.New("list size less than 1"), err)

	a = 10
	st, err = InitArrayStack(a)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, st.Length())
	assert.Equal(t, a, st.Size())
}

func Test_GetTop(t *testing.T) {
	a := 10
	st, _ := InitArrayStack(a)
	st.Push(1)
	elem := st.GetTop()
	assert.Equal(t, 1, elem.(int))
}

func Test_Push(t *testing.T) {
	a := 2
	st, _ := InitArrayStack(a)
	st.Push(3)
	elem := st.GetTop()
	assert.Equal(t, 3, elem.(int))
	st.Push(4)
	elem = st.GetTop()
	assert.Equal(t, 4, elem.(int))

}

func Test_Pop(t *testing.T) {
	a := 2
	st, _ := InitArrayStack(a)
	st.Push(3)
	elem := st.Pop()
	assert.Equal(t, 3, elem.(int))
	st.Push(4)
	elem = st.Pop()
	assert.Equal(t, 4, elem.(int))
	elem = st.Pop()
	assert.Equal(t, nil, elem)
}

func Test_ClearStack(t *testing.T) {

}
