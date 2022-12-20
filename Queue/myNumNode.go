package myNode

type MyNode struct {
	value    int
	nextNode *MyNode
}

func (n *MyNode) GetValue() int {
	return n.value
}

func (n *MyNode) GetNextNode() *MyNode {
	return n.nextNode
}
