package TreeNode

import (
	"errors"
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

func (t *TreeNode) GetChildren() []*TreeNode {
	return t.children
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

// BreadthFirstSearch will search target from most parent node, level to level.
// so the bottom level will be searched last
func BreadthFirstSearch(tree *TreeNode, target int) ([]TreeNode, error) {
	frontier := [][]TreeNode{}
	initPath := [][]TreeNode{{*tree}}

	frontier = append(initPath, frontier...)

	for len(frontier) > 0 {
		//pop the last, slice shortened.
		currentPath := frontier[len(frontier)-1]
		frontier = frontier[0 : len(frontier)-1]

		currentNode := currentPath[len(currentPath)-1]
		fmt.Printf("Searching node with value: %d \n", currentNode.value)

		if currentNode.value == target {
			return currentPath, nil
		}

		// current node not target,
		// adding the children nodes to frontier path.
		for _, child := range currentNode.children {
			newPath := make([]TreeNode, len(currentPath))

			// for copy function, new slice as first argument,
			// original slice as second argument
			copy(newPath, currentPath)
			newPath = append(newPath, *child)

			// new children are added at front,
			// as those at the back of slice will be popped,
			// these new children will be searched LATER than this level
			frontier = append([][]TreeNode{newPath}, frontier...)
		}
	}
	return nil, errors.New("not found")
}

func DepthFirstSearch(root *TreeNode, target int) ([]*TreeNode, error) {
	fmt.Println("finding ", target)
	resultPath := Dfs(root, target, []*TreeNode{})
	if resultPath == nil {
		return nil, errors.New("target not found")
	} else {
		return resultPath, nil
	}
}

func Dfs(root *TreeNode, target int, path []*TreeNode) []*TreeNode {

	fmt.Println("at node: ", root.value)
	path = append(path, root)

	if root.value == target {
		return path
	}

	for _, child := range root.children {
		// this error is taken up by _
		// because it only means target does not appear in this tree
		pathFound := Dfs(child, target, path)

		if len(pathFound) > 0 {
			return pathFound
		}
	}

	return nil
}
