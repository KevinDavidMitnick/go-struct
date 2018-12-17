package doublelinklist

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InitLinkList(t *testing.T) {
	lst, err := InitLinkList()
	assert.NotEqual(t, nil, lst)
	assert.Equal(t, nil, err)

	assert.Equal(t, 0, lst.(*LinkList).Length())
	assert.Equal(t, 0, lst.(*LinkList).Size())
}

func Test_IsEmpty(t *testing.T) {
	lst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, lst.IsEmpty())
	err = lst.Insert(0, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, lst.Length())
	assert.Equal(t, false, lst.IsEmpty())
}

func Test_Length(t *testing.T) {
	a := 4
	lst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, lst.Length())
	lst.Insert(0, 1)
	lst.Insert(0, 1)
	lst.Insert(0, 1)
	lst.Insert(0, 1)
	assert.Equal(t, a, lst.Length())
}

func Test_Size(t *testing.T) {
	a := 2
	lst, err := InitLinkList()
	assert.Equal(t, nil, err)
	lst.Insert(0, 1)
	lst.Insert(0, 1)
	assert.Equal(t, a, lst.Size())
}

func Test_IsFull(t *testing.T) {
	lst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, false, lst.IsFull())
	lst.Insert(0, 1)
	lst.Insert(0, 1)
	lst.Insert(0, 1)
	lst.Insert(0, 1)
	assert.Equal(t, false, lst.IsFull())
}

func Test_ClearList(t *testing.T) {
	lst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, lst.IsEmpty())
	lst.Insert(0, 1)
	lst.Insert(0, 1)
	lst.Insert(0, 1)
	lst.Insert(0, 1)
	assert.Equal(t, false, lst.IsEmpty())
	lst.ClearList()
	assert.Equal(t, true, lst.IsEmpty())
}

func Test_GetElem(t *testing.T) {
	lst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, lst.IsEmpty())

	item, e := lst.GetElem(0)
	assert.Equal(t, nil, item)
	assert.Equal(t, errors.New("list is empty."), e)

	lst.Insert(0, 1)
	item, e = lst.GetElem(0)
	assert.Equal(t, 1, item.(int))
	assert.Equal(t, nil, e)

	lst.Insert(0, 2)
	item, e = lst.GetElem(0)
	assert.Equal(t, 2, item.(int))
	assert.Equal(t, nil, e)

	lst.Insert(1, 3)
	item, e = lst.GetElem(1)
	assert.Equal(t, 3, item.(int))
	assert.Equal(t, nil, e)

	lst.Insert(3, 4)
	item, e = lst.GetElem(3)
	assert.Equal(t, 4, item.(int))
	assert.Equal(t, nil, e)

	item, e = lst.GetElem(5)
	assert.Equal(t, nil, item)
	assert.Equal(t, errors.New("index i not valid."), e)
}

func Test_LocateElem(t *testing.T) {
	lst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, lst.IsEmpty())

	item := lst.LocateElem(23)
	assert.Equal(t, -1, item)

	lst.Insert(0, 1)
	item = lst.LocateElem(1)
	assert.Equal(t, 0, item)

	lst.Insert(0, 2)
	item = lst.LocateElem(2)
	assert.Equal(t, 0, item)

	lst.Insert(1, 3)
	item = lst.LocateElem(3)
	assert.Equal(t, 1, item)

	lst.Insert(3, 4)
	item = lst.LocateElem(4)
	assert.Equal(t, 3, item)

	item = lst.LocateElem(1)
	assert.Equal(t, 2, item)
}

func Test_Delete(t *testing.T) {
	lst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, lst.IsEmpty())

	lst.Insert(0, 1)
	lst.Insert(1, 2)
	lst.Insert(2, 3)
	lst.Insert(3, 4)

	item, e := lst.Delete(0)
	assert.Equal(t, nil, e)
	assert.Equal(t, 1, item.(int))

	item, e = lst.Delete(0)
	assert.Equal(t, nil, e)
	assert.Equal(t, 2, item.(int))

	item, e = lst.Delete(2)
	assert.Equal(t, errors.New("index i not valid."), e)
	assert.Equal(t, nil, item)

	item, e = lst.Delete(1)
	assert.Equal(t, nil, e)
	assert.Equal(t, 4, item.(int))

	item, e = lst.Delete(0)
	assert.Equal(t, nil, e)
	assert.Equal(t, 3, item.(int))

	item, e = lst.Delete(0)
	assert.Equal(t, errors.New("list is empty."), e)
	assert.Equal(t, nil, item)
}

func Test_Equal(t *testing.T) {
	lst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, lst.IsEmpty())

	lst.Insert(0, 1)
	lst.Insert(1, 2)
	lst.Insert(2, 3)
	lst.Insert(3, 4)

	olst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, olst.IsEmpty())

	olst.Insert(0, 1)
	olst.Insert(1, 2)
	olst.Insert(2, 3)
	olst.Insert(3, 4)

	blst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, blst.IsEmpty())

	blst.Insert(0, 1)
	blst.Insert(1, 2)
	blst.Insert(2, 3)
	blst.Insert(3, 5)

	flag1 := lst.Equal(olst)
	flag2 := lst.Equal(blst)
	assert.Equal(t, true, flag1)
	assert.Equal(t, false, flag2)
}

func Test_Clone(t *testing.T) {
	lst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, lst.IsEmpty())

	lst.Insert(0, 1)
	lst.Insert(1, 2)
	lst.Insert(2, 3)
	lst.Insert(3, 4)

	olst := lst.Clone()
	flag1 := lst.Equal(olst)
	assert.Equal(t, true, flag1)
}

func Test_Diff(t *testing.T) {
	lst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, lst.IsEmpty())

	lst.Insert(0, 1)
	lst.Insert(1, 2)
	lst.Insert(2, 3)
	lst.Insert(3, 4)

	olst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, olst.IsEmpty())
	olst.Insert(0, 2)
	olst.Insert(1, 3)

	clst := lst.Diff(olst)

	assert.Equal(t, 2, clst.Length())
	assert.Equal(t, 2, clst.Size())
	item1, _ := clst.GetElem(0)
	item2, _ := clst.GetElem(1)
	assert.Equal(t, 1, item1)
	assert.Equal(t, 4, item2)
}

func Test_Union(t *testing.T) {
	lst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, lst.IsEmpty())

	lst.Insert(0, 1)
	lst.Insert(1, 3)

	olst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, olst.IsEmpty())

	olst.Insert(0, 1)
	olst.Insert(1, 2)

	blst, err := InitLinkList()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, blst.IsEmpty())

	blst.Insert(0, 1)
	blst.Insert(1, 3)
	blst.Insert(2, 2)

	dlst := lst.Union(olst)
	assert.Equal(t, true, dlst.Equal(blst))
}
