package Queue

import "errors"

type MyQueue struct {
	head *MyNode
	cap  int
	size int
}

func (q *MyQueue) GetCap() int {
	return q.cap
}

func (q *MyQueue) CurrentSize() int {
	return q.size
}

func (q *MyQueue) Peek() (n int, err error) {
	if q.head != nil {
		return q.head.value, nil
	} else {
		return -1, errors.New("the queue is empty")
	}
}

func (q *MyQueue) Enqueue(value int) error {
	if q.cap > q.size {
		newNode := MyNode{value: value}
		if q.head != nil {
			q.head.SetNext(newNode)
		} else {
			q.head = &newNode
		}
		q.size += 1

		return nil
	} else {
		return errors.New("max capacity reached, cannot enqueue new data")
	}
}

func (q *MyQueue) Dequeue() (int, error) {
	if q.size == 0 {
		return -1, errors.New("the queue is empty, dequeue will cause stack underflow")
	}
	if q.head.nextNode == nil {
		removed := q.head
		q.head = nil
		q.size -= 1
		return removed.GetValue(), nil
	} else {
		removed := q.head
		q.head = q.head.nextNode
		q.size -= 1
		return removed.GetValue(), nil
	}
}

func (q *MyQueue) GetWholeList() string {
	return ""
}
