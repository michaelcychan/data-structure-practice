package TreeNode

import "testing"

func TestTreeNode(t *testing.T) {
	t.Run("create a TreeNode", func(t *testing.T) {
		newTreeNode := TreeNode{value: 100}
		gotValue := newTreeNode.GetValue()

		if gotValue != 100 {
			t.Errorf("expected the value to be %d, but got %d", 100, gotValue)
		}
	})
	t.Run("creating a TreeNode using CreateTreeNode()", func(t *testing.T) {
		newTreeNode := CreateTreeNode(100)
		gotValue := newTreeNode.GetValue()

		if gotValue != 100 {
			t.Errorf("expected the value to be %d, but got %d", 100, gotValue)
		}
	})
	t.Run("change treenode value", func(t *testing.T) {
		newTreeNode := CreateTreeNode(100)

		newTreeNode.SetValue(105)
		gotValue := newTreeNode.GetValue()
		if gotValue != 105 {
			t.Errorf("expected the value to be %d, but got %d", 100, gotValue)
		}
	})

}
