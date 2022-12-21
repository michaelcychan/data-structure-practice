package Queue

type MyNode struct {
	value    int
	nextNode *MyNode
}

func (n *MyNode) GetValue() int {
	return n.value
}

func (n *MyNode) GetNext() *MyNode {
	return n.nextNode
}

func (n *MyNode) SetNext(nextNode MyNode) {
	n.nextNode = &nextNode
}

func (n *MyNode) RemoveNext() {
	n.nextNode = nil
}
