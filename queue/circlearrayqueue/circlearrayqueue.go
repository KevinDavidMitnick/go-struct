package circlearrayqueue

import (
	"errors"
	"github.com/go-struct/queue"
)

type CircleArrayQueue struct {
	front int
	rear  int
	Item  []interface{}
}

func InitCircleArrayQueue(size int) (queue.Queue, error) {
	if size < 2 {
		return nil, errors.New("list size less than 2")
	}
	var circleArrayQueue CircleArrayQueue
	circleArrayQueue.Item = make([]interface{}, size)
	circleArrayQueue.front = 0
	circleArrayQueue.rear = 0
	return &circleArrayQueue, nil
}

func (queue *CircleArrayQueue) IsFull() bool {
	if queue.Length()+1 >= queue.Size() {
		return true
	}
	return false
}

func (queue *CircleArrayQueue) Size() int {
	return len(queue.Item)
}

func (queue *CircleArrayQueue) Length() int {
	if queue.rear-queue.front >= 0 {
		return queue.rear - queue.front
	}
	return queue.rear - queue.front + queue.Size()
}

func (queue *CircleArrayQueue) IsEmpty() bool {
	if queue.front == queue.rear {
		return true
	}
	return false
}

func (queue *CircleArrayQueue) ClearQueue() {
	queue.front = 0
	queue.rear = 0
}

func (queue *CircleArrayQueue) GetHead() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	elem := queue.Item[queue.front]
	return elem
}

func (queue *CircleArrayQueue) EnQueue(elem interface{}) {
	if queue.IsFull() {
		return
	}
	queue.Item[queue.rear] = elem
	queue.rear = queue.rear%queue.Size() + 1
}

func (queue *CircleArrayQueue) DeQueue() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	elem := queue.Item[queue.front]
	queue.front = queue.front%queue.Size() + 1
	return elem
}
