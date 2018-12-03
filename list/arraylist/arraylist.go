package arraylist

import (
	"errors"
	"fmt"
	"github.com/KevinDavidMitnick/go-struct/list"
)

type ArrayList struct {
	len  int
	Item []interface{}
}

func InitArrayList(size int) (list.List, error) {
	var arrayList ArrayList
	if size < 1 {
		return nil, errors.New("list size less than 1")
	}
	arrayList.len = 0
	arrayList.Item = make([]interface{}, size)
	return &arrayList, nil
}

func (l *ArrayList) IsEmpty() bool {
	return 0 == l.len
}

func (l *ArrayList) Length() int {
	return l.len
}

func (l *ArrayList) Size() int {
	return len(l.Item)
}

func (l *ArrayList) IsFull() bool {
	if l.IsEmpty() {
		return false
	}
	if l.len == len(l.Item) {
		return true
	}
	return false
}

func (l *ArrayList) ClearList() {
	l.len = 0
}

func (l *ArrayList) GetElem(i int) (interface{}, error) {
	if l.IsEmpty() {
		return nil, errors.New("list is empty.")
	}
	if i < 0 || i >= l.len {
		return nil, errors.New("index i not valid.")
	}
	return l.Item[i], nil
}

func (l *ArrayList) LocateElem(e interface{}) int {
	for i := 0; i < l.len; i++ {
		if l.Item[i] == e {
			return i
		}
	}
	return -1
}

func (l *ArrayList) Insert(i int, e interface{}) error {
	if l.IsFull() {
		return errors.New("list is full.")
	}
	if i < 0 || i > l.len {
		return errors.New("index i not valid.")
	}
	if i < l.len {
		for j := l.len; j > i; j-- {
			l.Item[j] = l.Item[j-1]
		}
	}
	l.Item[i] = e
	l.len += 1
	return nil
}

func (l *ArrayList) Delete(i int) (interface{}, error) {
	if l.IsEmpty() {
		return nil, errors.New("list is empty.")
	}
	if i < 0 || i >= l.len {
		return nil, errors.New("index i not valid.")
	}
	hold := l.Item[i]
	for j := i; j < l.len-1; j++ {
		l.Item[j] = l.Item[j+1]
	}
	l.len -= 1
	return hold, nil
}

func (l *ArrayList) Equal(list list.List) bool {
	ol := list.(*ArrayList)
	if l.Length() != ol.Length() {
		return false
	}
	if l.Size() != ol.Size() {
		return false
	}
	for i := 0; i < l.Length(); i++ {
		Item1, _ := l.GetElem(i)
		Item2, _ := list.GetElem(i)
		if Item1 != Item2 {
			return false
		}
	}
	return true
}

func (l *ArrayList) Clone() list.List {
	list, _ := InitArrayList(l.Size())
	nl := list.(*ArrayList)
	for i := 0; i < l.len; i++ {
		Item, _ := l.GetElem(i)
		nl.Insert(i, Item)
	}
	return list
}

func (l *ArrayList) Diff(list list.List) list.List {
	if list == nil || list.IsEmpty() || l.IsEmpty() {
		return l.Clone()
	}
	ol := l.Clone()
	al := list.(*ArrayList)
	for i := 0; i < al.len; i++ {
		if j := ol.LocateElem(al.Item[i]); j >= 0 {
			ol.Delete(j)
		}
	}
	dl, _ := InitArrayList(ol.Length())
	for i := 0; i < ol.Length(); i++ {
		Item, _ := ol.GetElem(i)
		dl.Insert(dl.Length(), Item)
	}
	return dl
}

func (l *ArrayList) Union(list list.List) list.List {
	if list == nil || list.IsEmpty() {
		return l.Clone()
	}
	if l.IsEmpty() {
		return list.Clone()
	}
	ol, _ := InitArrayList(l.Length() + list.Length())
	for i := 0; i < l.len; i++ {
		ol.Insert(i, l.Item[i])
	}
	al := list.(*ArrayList)
	for i := 0; i < al.Length(); i++ {
		Item, _ := al.GetElem(i)
		if ol.LocateElem(Item) < 0 {
			ol.Insert(ol.Length(), Item)
		}
	}
	dl, _ := InitArrayList(ol.Length())
	for i := 0; i < ol.Length(); i++ {
		Item, _ := ol.GetElem(i)
		dl.Insert(dl.Length(), Item)
	}

	return dl
}

func (l *ArrayList) Display() {
	for i := 0; i < l.len; i++ {
		fmt.Print(" ", l.Item[i])
	}
	fmt.Println()
}
