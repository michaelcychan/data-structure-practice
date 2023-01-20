package TreeNode

import "testing"

func TestTreeNode(t *testing.T) {
	newTreeNode := TreeNode{value: 100}
	gotValue := newTreeNode.GetValue()

	if gotValue != 100 {
		t.Errorf("expected the value to be %d, but got %d", 100, gotValue)
	}
}
