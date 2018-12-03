package linkstack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InitLinkStack(t *testing.T) {
	st, err := InitLinkStack()
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, st)
	assert.Equal(t, 0, st.Length())
	assert.Equal(t, 0, st.Size())
}

func Test_GetTop(t *testing.T) {
	st, _ := InitLinkStack()
	st.Push(1)
	elem := st.GetTop()
	assert.Equal(t, 1, elem.(int))
}

func Test_Push(t *testing.T) {
	st, _ := InitLinkStack()
	st.Push(3)
	elem := st.GetTop()
	assert.Equal(t, 3, elem.(int))
	st.Push(4)
	elem = st.GetTop()
	assert.Equal(t, 4, elem.(int))

}

func Test_Pop(t *testing.T) {
	st, _ := InitLinkStack()
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
