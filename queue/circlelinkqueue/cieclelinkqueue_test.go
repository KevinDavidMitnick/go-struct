package circlelinkqueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InitCircleLinkQueue(t *testing.T) {
	st, err := InitCircleLinkQueue()
	assert.NotEqual(t, nil, st)
	assert.Equal(t, nil, err)

	st, err = InitCircleLinkQueue()
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, st.Length())
	assert.Equal(t, 0, st.Size())
}

func Test_GetHead(t *testing.T) {
	st, _ := InitCircleLinkQueue()
	st.EnQueue(1)
	elem := st.GetHead()
	assert.Equal(t, 1, elem.(int))
	st.EnQueue(3)
	elem = st.GetHead()
	assert.Equal(t, 1, elem.(int))
}

func Test_EnQueue(t *testing.T) {
	st, _ := InitCircleLinkQueue()
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
	st, _ := InitCircleLinkQueue()
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
	st, _ := InitCircleLinkQueue()
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
