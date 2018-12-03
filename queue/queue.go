package queue

type Queue interface {
	ClearQueue()
	GetHead() interface{}
	Length() int
	Size() int
	IsFull() bool
	IsEmpty() bool
	EnQueue(elem interface{})
	DeQueue() interface{}
}
