package linkstack

import (
	"github.com/KevinDavidMitnick/go-struct/list/linklist"
	"github.com/KevinDavidMitnick/go-struct/stack"
)

type LinkStack struct {
	linklist.LinkList
}

func InitLinkStack() (stack.Stack, error) {
	var linkStack LinkStack
	linkStack.Item = nil
	linkStack.Next = &linklist.Node{Item: nil, Next: nil}
	linkStack.Len = 0
	return &linkStack, nil
}

func (stack *LinkStack) GetTop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.Next.Next.Item
}

func (stack *LinkStack) Push(elem interface{}) {
	stack.Insert(0, elem)
}

func (stack *LinkStack) Pop() interface{} {
	elem, _ := stack.Delete(0)
	return elem
}

func (stack *LinkStack) ClearStack() {
	stack.ClearList()
}
