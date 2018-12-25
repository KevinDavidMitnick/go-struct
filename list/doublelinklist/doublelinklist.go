package doublelinklist

import (
	"errors"
	"fmt"
	"github.com/KevinDavidMitnick/go-struct/list"
)

type Node struct {
	Item interface{}
	Next *Node
	Pre  *Node
}

type LinkList struct {
	Node
	Len  int
	Tail *Node
}

func InitLinkList() (list.List, error) {
	var linkList LinkList
	linkList.Item = nil
	node := Node{Item: nil, Next: nil, Pre: nil}
	node.Pre = &node
	node.Next = &node
	linkList.Next = &node
	linkList.Tail = &node
	linkList.Len = 0
	return &linkList, nil
}

func (l *LinkList) IsEmpty() bool {
	if l.Len == 0 {
		return true
	}
	return false
}

func (l *LinkList) Length() int {
	return l.Len
}

func (l *LinkList) Size() int {
	return l.Len
}

func (l *LinkList) Equal(ol list.List) bool {
	if l.Len != ol.Length() {
		return false
	}
	p := l.Next.Next
	o := ol.(*LinkList)
	q := o.Next.Next
	for p != l.Next && q != o.Next {
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
	if i < 0 || i >= l.Len {
		return nil, errors.New("index i not valid.")
	}
	p := l.Next.Next
	j := 0
	for p != l.Next && j < i {
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
	for p != l.Next {
		if p.Item == e {
			return j
		}
		j += 1
		p = p.Next
	}
	return -1
}

func (l *LinkList) Insert(i int, e interface{}) error {
	if i < 0 || i > l.Len {
		return errors.New("index i not valid.")
	}

	if i == l.Len {
		t := &Node{Item: e, Next: l.Tail.Next, Pre: l.Tail}
		l.Tail.Next.Pre = t
		l.Tail.Next = t
		l.Tail = t
		l.Len += 1
		return nil
	}

	p := l.Next
	j := 0
	for j < i {
		p = p.Next
		j += 1
	}
	node := &Node{Item: e, Next: p.Next, Pre: p}
	p.Next.Pre = node
	p.Next = node
	l.Len += 1
	return nil
}

func (l *LinkList) Delete(i int) (interface{}, error) {
	if l.IsEmpty() {
		return nil, errors.New("list is empty.")
	}
	if i < 0 || i >= l.Len {
		return nil, errors.New("index i not valid.")
	}
	p := l.Next
	j := 0
	for p.Next != l.Next && j < i {
		p = p.Next
		j += 1
	}

	p.Next.Next.Pre = p
	item := p.Next.Item
	p.Next = p.Next.Next
	l.Len -= 1
	return item, nil
}

func (l *LinkList) Clone() list.List {
	var linkList LinkList
	linkList.Item = nil
	node := Node{Item: nil, Next: nil, Pre: nil}
	node.Pre = &node
	node.Next = &node
	linkList.Next = &node
	linkList.Tail = &node
	linkList.Len = 0

	p := l.Next.Next
	for i := 0; i < l.Len; i++ {
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
	for i := 0; p != nil && i < al.Len; i++ {
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
	for i := 0; p != nil && i < al.Len; i++ {
		if ol.LocateElem(p.Item) < 0 {
			ol.Insert(ol.Length(), p.Item)
		}
		p = p.Next
	}
	return ol
}

func (l *LinkList) ClearList() {
	l.Next.Pre = l.Next
	l.Next.Next = l.Next
	l.Tail = l.Next
	l.Len = 0
}

func (l *LinkList) Display() {
	p := l.Next.Next
	for ; p != l.Next; p = p.Next {
		fmt.Print(" ", p.Item)
	}
	fmt.Println()
}
