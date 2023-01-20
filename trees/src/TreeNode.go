package TreeNode

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	value    int
	children []TreeNode
}

type NodeLevel struct {
	node  TreeNode
	level int
}

func (t *TreeNode) GetValue() int {
	return t.value
}

func (t *TreeNode) AddChild(child TreeNode) {
	t.children = append(t.children, child)
}

func PrintTree(t TreeNode) {
	stack := []NodeLevel{{t, 0}}

	// collecting strings
	var sb strings.Builder
	sb.WriteString("\n")

	for len(stack) > 0 {
		currentLength := len(stack)
		node := stack[currentLength-1]
		level := node.level
		stack = stack[0 : currentLength-1]

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
			stack = append(stack, NodeLevel{child, level})

		}
	}
}
