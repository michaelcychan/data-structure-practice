package TreeNode

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	value int

	// children should saves the pointers,
	// so it shows the original copy,
	// not just a shadow copy
	children []*TreeNode
}

type NodeLevel struct {
	node  TreeNode
	level int
}

func (t *TreeNode) GetValue() int {
	return t.value
}

func (t *TreeNode) SetValue(value int) {
	t.value = value
}

func (t *TreeNode) AddChild(child *TreeNode) {
	t.children = append(t.children, child)
}

func CreateTreeNode(value int) TreeNode {
	return TreeNode{value: value}
}

func PrintTree(t TreeNode) {
	fmt.Println("Printing the tree")
	stack := []NodeLevel{{t, 0}}

	// collecting strings
	var sb strings.Builder
	sb.WriteString("\n")

	var node NodeLevel
	var level int

	for len(stack) > 0 {
		currentLength := len(stack)
		node = stack[0]
		level = node.level
		stack = stack[1:currentLength]

		// write the level indicators
		if level > 0 {
			for i := 1; i < level; i++ {
				sb.WriteString("| ")
			}
			sb.WriteString("|-")
		}

		// write exact value
		sb.WriteString(strconv.Itoa(node.node.value))
		sb.WriteString("\n")

		level++

		for _, child := range node.node.children {
			stack = append([]NodeLevel{{*child, level}}, stack...)
		}
	}
	fmt.Println(sb.String())
}
