package MyStack

import (
	"Queue/src/MyNode"
	"errors"
)

type MyStack struct {
	head *MyNode.MyNode
	cap  int
	size int
}

func (s *MyStack) Peek() (n int, err error) {
	if s.head == nil {
		return -1, errors.New("stack is empty")
	} else {
		return s.head.GetValue(), nil
	}
}

func (s *MyStack) Push(value int) error {
	if s.size < s.cap {
		newNode := MyNode.MyNode{}
		newNode.SetValue(value)
		if s.head != nil {
			newNode.SetNext(*s.head)
		}

		s.head = &newNode

		s.size++
		return nil
	} else {
		return errors.New("cap reached, stack overflow")
	}
}

func (s *MyStack) Pop() (topValue int, err error) {
	if s.size == 0 {
		return -1, errors.New("stack underflow")
	} else {
		currentTop := s.head
		s.head = currentTop.GetNext()
		s.size--
		return currentTop.GetValue(), nil
	}
}
