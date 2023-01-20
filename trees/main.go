package main

import (
	TreeNode "github.com/michaelcychan/data-structure-practice/trees/src"
)

func main() {

	newTreeNode := TreeNode.CreateTreeNode(100)
	level1Child1 := TreeNode.CreateTreeNode(80)
	level1Child2 := TreeNode.CreateTreeNode(120)

	level2Child11 := TreeNode.CreateTreeNode(50)
	level2Child12 := TreeNode.CreateTreeNode(60)
	level2Child13 := TreeNode.CreateTreeNode(75)

	level2Child21 := TreeNode.CreateTreeNode(175)

	newTreeNode.AddChild(&level1Child1)
	newTreeNode.AddChild(&level1Child2)

	level1Child2.AddChild(&level2Child21)

	level1Child1.AddChild(&level2Child11)
	level1Child1.AddChild(&level2Child12)
	level1Child1.AddChild(&level2Child13)

	newTreeNode.SetValue(101)

	level3Child31 := TreeNode.CreateTreeNode(300)
	level2Child12.AddChild(&level3Child31)

	TreeNode.PrintTree(newTreeNode)
}
