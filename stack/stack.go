package stack

type Stack interface {
	IsEmpty() bool
	IsFull() bool
	Size() int
	Length() int
	GetTop() interface{}
	Push(elem interface{})
	Pop() interface{}
	ClearStack()
}
