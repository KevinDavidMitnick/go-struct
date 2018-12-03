package arrayqueue

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InitArrayQueue(t *testing.T) {
	a := -1
	st, err := InitArrayQueue(a)
	assert.Equal(t, nil, st)
	assert.Equal(t, errors.New("list size less than 1"), err)

	a = 0
	st, err = InitArrayQueue(a)
	assert.Equal(t, nil, st)
	assert.Equal(t, errors.New("list size less than 1"), err)

	a = 10
	st, err = InitArrayQueue(a)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, st.Length())
	assert.Equal(t, a, st.Size())
}

func Test_GetHead(t *testing.T) {
	a := 10
	st, _ := InitArrayQueue(a)
	st.EnQueue(1)
	elem := st.GetHead()
	assert.Equal(t, 1, elem.(int))
	st.EnQueue(3)
	elem = st.GetHead()
	assert.Equal(t, 1, elem.(int))
}

func Test_EnQueue(t *testing.T) {
	a := 10
	st, _ := InitArrayQueue(a)
	elem := st.GetHead()
	assert.Equal(t, nil, elem)
	st.EnQueue(1)
	elem = st.GetHead()
	assert.Equal(t, 1, elem.(int))
	st.EnQueue(3)
	elem = st.GetHead()
	assert.Equal(t, 1, elem.(int))
}

func Test_DeQueu(t *testing.T) {
	a := 10
	st, _ := InitArrayQueue(a)
	elem := st.GetHead()
	assert.Equal(t, nil, elem)
	st.EnQueue(1)
	st.EnQueue(3)
	elem = st.DeQueue()
	assert.Equal(t, 1, elem.(int))
	elem = st.DeQueue()
	assert.Equal(t, 3, elem.(int))
	elem = st.DeQueue()
	assert.Equal(t, nil, elem)
}

func Test_ClearQueu(t *testing.T) {
	a := 10
	st, _ := InitArrayQueue(a)
	elem := st.GetHead()
	assert.Equal(t, nil, elem)
	st.EnQueue(1)
	st.EnQueue(3)
	st.EnQueue(4)
	st.EnQueue(3)
	assert.Equal(t, false, st.IsEmpty())
	st.ClearQueue()
	assert.Equal(t, true, st.IsEmpty())
}
