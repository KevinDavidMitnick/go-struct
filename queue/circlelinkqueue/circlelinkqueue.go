package circlelinkqueue

import (
	"github.com/KevinDavidMitnick/go-struct/list/circlelinklist"
	"github.com/KevinDavidMitnick/go-struct/queue"
)

type CircleLinkQueue struct {
	circlelinklist.CircleLinkList
}

func InitCircleLinkQueue() (queue.Queue, error) {
	var circleLinkQueue CircleLinkQueue
	circleLinkQueue.Item = nil
	circleLinkQueue.Next = &circlelinklist.Node{Item: nil, Next: nil}
	circleLinkQueue.Next.Next = circleLinkQueue.Next
	circleLinkQueue.Tail = circleLinkQueue.Next
	circleLinkQueue.Len = 0
	return &circleLinkQueue, nil
}

func (queue *CircleLinkQueue) ClearQueue() {
	queue.ClearList()
}

func (queue *CircleLinkQueue) GetHead() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	head, _ := queue.GetElem(0)
	return head
}

func (queue *CircleLinkQueue) EnQueue(elem interface{}) {
	if queue.IsFull() {
		return
	}
	queue.Insert(queue.Length(), elem)
}

func (queue *CircleLinkQueue) DeQueue() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	elem := queue.GetHead()
	queue.Delete(0)
	return elem
}
