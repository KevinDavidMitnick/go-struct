package arraystack

import (
	"errors"
	"github.com/KevinDavidMitnick/go-struct/list/arraylist"
	"github.com/KevinDavidMitnick/go-struct/stack"
)

type ArrayStack struct {
	arraylist.ArrayList
}

func InitArrayStack(size int) (stack.Stack, error) {
	if size < 1 {
		return nil, errors.New("list size less than 1")
	}
	var arrayStack ArrayStack
	arrayStack.Item = make([]interface{}, size)
	return &arrayStack, nil
}

func (stack *ArrayStack) GetTop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	elem, _ := stack.GetElem(stack.Length() - 1)
	return elem
}

func (stack *ArrayStack) Push(elem interface{}) {
	stack.Insert(stack.Length(), elem)
}

func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	elem, err := stack.GetElem(stack.Length() - 1)
	if err == nil {
		stack.Delete(stack.Length() - 1)
	}
	return elem
}

func (stack *ArrayStack) ClearStack() {
	stack.ClearList()
}
