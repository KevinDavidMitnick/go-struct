package arrayqueue

import (
	"errors"
	"github.com/KevinDavidMitnick/go-struct/list/arraylist"
	"github.com/KevinDavidMitnick/go-struct/queue"
)

type ArrayQueue struct {
	arraylist.ArrayList
}

func InitArrayQueue(size int) (queue.Queue, error) {
	if size < 1 {
		return nil, errors.New("list size less than 1")
	}
	var arrayQueue ArrayQueue
	arrayQueue.Item = make([]interface{}, size)
	return &arrayQueue, nil
}

func (queue *ArrayQueue) ClearQueue() {
	queue.ClearList()
}

func (queue *ArrayQueue) GetHead() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	elem, _ := queue.GetElem(0)
	return elem
}

func (queue *ArrayQueue) EnQueue(elem interface{}) {
	queue.Insert(queue.Length(), elem)
}

func (queue *ArrayQueue) DeQueue() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	elem, err := queue.GetElem(0)
	if err == nil {
		queue.Delete(0)
	}
	return elem
}
