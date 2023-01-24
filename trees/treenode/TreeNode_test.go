package TreeNode

import (
	"fmt"
	"reflect"
	"testing"
)

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

func TestBreadthFirstSearch(t *testing.T) {
	root := CreateTreeNode(100)
	child1 := CreateTreeNode(80)
	child2 := CreateTreeNode(120)

	grandchild1 := CreateTreeNode(3)
	grandchild2 := CreateTreeNode(220)
	grandchild3 := CreateTreeNode(240)
	grandchild4 := CreateTreeNode(255)

	greatgrandchild1 := CreateTreeNode(1050)
	greatgrandchild2 := CreateTreeNode(980)

	root.AddChild(&child1)
	root.AddChild(&child2)

	child1.AddChild(&grandchild1)

	child2.AddChild(&grandchild2)
	child2.AddChild(&grandchild3)
	child2.AddChild(&grandchild4)

	grandchild3.AddChild(&greatgrandchild1)
	grandchild3.AddChild(&greatgrandchild2)

	cases := []struct {
		target       int
		expectedPath []TreeNode
	}{
		{100, []TreeNode{root}},
		{80, []TreeNode{root, child1}},
		{240, []TreeNode{root, child2, grandchild3}},
		{1050, []TreeNode{root, child2, grandchild3, greatgrandchild1}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("searching for %d", c.target), func(t *testing.T) {
			gotPath, err := BreadthFirstSearch(&root, c.target)

			assertNoError(t, err)

			if !reflect.DeepEqual(gotPath, c.expectedPath) {
				t.Errorf("expected path to be %v, but got %v", c.expectedPath, gotPath)
			}
		})
	}
	t.Run("should returns if target not found", func(t *testing.T) {
		_, err := BreadthFirstSearch(&root, 507)

		if err == nil {
			t.Errorf("expected an error, but got none")
		}
	})
}

func TestDepthFirstSearch(t *testing.T) {
	root := CreateTreeNode(100)
	child1 := CreateTreeNode(80)
	child2 := CreateTreeNode(120)

	grandchild1 := CreateTreeNode(3)
	grandchild2 := CreateTreeNode(220)
	grandchild3 := CreateTreeNode(240)
	grandchild4 := CreateTreeNode(255)

	greatgrandchild1 := CreateTreeNode(1050)
	greatgrandchild2 := CreateTreeNode(980)

	root.AddChild(&child1)
	root.AddChild(&child2)

	child1.AddChild(&grandchild1)

	child2.AddChild(&grandchild2)
	child2.AddChild(&grandchild3)
	child2.AddChild(&grandchild4)

	grandchild3.AddChild(&greatgrandchild1)
	grandchild3.AddChild(&greatgrandchild2)

	cases := []struct {
		target       int
		expectedPath []*TreeNode
	}{
		{100, []*TreeNode{&root}},
		{80, []*TreeNode{&root, &child1}},
		{240, []*TreeNode{&root, &child2, &grandchild3}},
		// {1050, []*TreeNode{&root, &child2, &grandchild3, &greatgrandchild1}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("searching for %d", c.target), func(t *testing.T) {
			gotPath, err := DepthFirstSearch(&root, c.target)

			assertNoError(t, err)

			if !reflect.DeepEqual(gotPath, c.expectedPath) {
				t.Errorf("expected path to be %v, but got %v", c.expectedPath, gotPath)
			}
		})
	}
	t.Run("should returns if target not found", func(t *testing.T) {
		_, err := DepthFirstSearch(&root, 507)

		if err == nil {
			t.Errorf("expected an error, but got none")
		}
	})
}

func assertNoError(t testing.TB, errorMsg error) {
	t.Helper()
	if errorMsg != nil {
		t.Errorf("expected no error, but got one: %v", errorMsg)
	}
}
