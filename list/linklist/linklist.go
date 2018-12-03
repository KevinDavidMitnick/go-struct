package linklist

import (
	"errors"
	"fmt"
	"github.com/KevinDavidMitnick/go-struct/list"
)

type Node struct {
	Item interface{}
	Next *Node
}

type LinkList struct {
	Node
	len int
}

func InitLinkList() (list.List, error) {
	var linkList LinkList
	linkList.Item = nil
	linkList.Next = &Node{Item: nil, Next: nil}
	linkList.len = 0
	return &linkList, nil
}

func (l *LinkList) IsEmpty() bool {
	if l.len == 0 {
		return true
	}
	return false
}

func (l *LinkList) Length() int {
	return l.len
}

func (l *LinkList) Size() int {
	return l.len
}

func (l *LinkList) Equal(ol list.List) bool {
	if l.len != ol.Length() {
		return false
	}
	p := l.Next
	q := ol.(*LinkList).Next
	for p != nil && q != nil {
		if p.Item != q.Item {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}

func (l *LinkList) IsFull() bool {
	return false
}

func (l *LinkList) GetElem(i int) (interface{}, error) {
	if l.IsEmpty() {
		return nil, errors.New("list is empty.")
	}
	if i < 0 || i >= l.len {
		return nil, errors.New("index i not valid.")
	}
	p := l.Next.Next
	j := 0
	for p != nil && j < i {
		p = p.Next
		j += 1
	}
	if j == i {
		return p.Item, nil
	}
	return nil, errors.New("not found,unknown error.")
}

func (l *LinkList) LocateElem(e interface{}) int {
	if l.IsEmpty() {
		return -1
	}
	p := l.Next.Next
	j := 0
	for p != nil {
		if p.Item == e {
			return j
		}
		j += 1
		p = p.Next
	}
	return -1
}

func (l *LinkList) Insert(i int, e interface{}) error {
	if i < 0 || i > l.len {
		return errors.New("index i not valid.")
	}
	p := l.Next
	j := 0
	for p.Next != nil && j < i {
		p = p.Next
		j += 1
	}
	p.Next = &Node{Item: e, Next: p.Next}
	l.len += 1
	return nil
}

func (l *LinkList) Delete(i int) (interface{}, error) {
	if l.IsEmpty() {
		return nil, errors.New("list is empty.")
	}
	if i < 0 || i >= l.len {
		return nil, errors.New("index i not valid.")
	}
	p := l.Next
	j := 0
	for p.Next != nil && j < i {
		p = p.Next
		j += 1
	}
	item := p.Next.Item
	p.Next = p.Next.Next
	l.len -= 1
	return item, nil
}

func (l *LinkList) Clone() list.List {
	var linkList LinkList
	linkList.Item = nil
	linkList.Next = &Node{Item: nil, Next: nil}
	linkList.len = 0

	p := l.Next.Next
	for i := 0; i < l.len; i++ {
		linkList.Insert(linkList.Length(), p.Item)
		p = p.Next
	}
	return &linkList
}

func (l *LinkList) Diff(list list.List) list.List {
	if list == nil || list.IsEmpty() || l.IsEmpty() {
		return l.Clone()
	}
	ol := l.Clone()
	al := list.(*LinkList)
	p := al.Next.Next
	for i := 0; p != nil && i < al.len; i++ {
		if j := ol.LocateElem(p.Item); j >= 0 {
			ol.Delete(j)
		}
		p = p.Next
	}
	return ol

}

func (l *LinkList) Union(list list.List) list.List {
	if list == nil || list.IsEmpty() {
		return l.Clone()
	}
	if l.IsEmpty() {
		return list.Clone()
	}

	ol := l.Clone()
	al := list.(*LinkList)

	p := al.Next.Next
	for i := 0; p != nil && i < al.len; i++ {
		if ol.LocateElem(p.Item) < 0 {
			ol.Insert(ol.Length(), p.Item)
		}
		p = p.Next
	}
	return ol
}

func (l *LinkList) ClearList() {
	l.Item = nil
	l.Next = nil
	l.len = 0
}

func (l *LinkList) Display() {
	p := l.Next
	for ; p != nil; p = p.Next {
		fmt.Print(" ", p.Item)
	}
	fmt.Println()
}
