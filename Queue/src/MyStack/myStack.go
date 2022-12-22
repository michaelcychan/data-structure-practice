package MyStack

import (
	"Queue/src/MyNode"
	"errors"
)

type MyStack struct {
	head *MyNode.MyNode
	cap  int
}

func (s *MyStack) Peek() (n int, err error) {
	if s.head == nil {
		return -1, errors.New("Stack is empty")
	} else {
		return s.head.GetValue(), nil
	}
}
