package BinaryTree

import (
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	t.Run("creates a binary tree", func(t *testing.T) {
		newNumber := 100
		treeroot := CreateBinaryTreeNode(newNumber)
		gotValue := treeroot.GetValue()

		assertValue(t, newNumber, gotValue)
	})
	t.Run("a new tree node has empty left and right node", func(t *testing.T) {
		treeroot := CreateBinaryTreeNode(100)
		left := treeroot.GetLeft()

		right := treeroot.GetRight()

		if left != nil {
			t.Errorf("left child should but empty, but got %v", left)
		}
		if right != nil {
			t.Errorf("right child should but empty, but got %v", left)
		}
	})
	t.Run("inserts a smaller number to root", func(t *testing.T) {
		treeroot := CreateBinaryTreeNode(100)
		newNumber := 90
		treeroot.Insert(newNumber)

		gotValue := treeroot.GetLeft().GetValue()

		assertValue(t, newNumber, gotValue)

	})
	t.Run("inserts a bigger number to root", func(t *testing.T) {
		treeroot := CreateBinaryTreeNode(100)
		newNumber := 120

		treeroot.Insert(newNumber)

		gotValue := treeroot.GetRight().GetValue()

		assertValue(t, newNumber, gotValue)
	})
	t.Run("inserts bigger numbers to root", func(t *testing.T) {
		treeroot := CreateBinaryTreeNode(100)
		newNumber := 120
		grandchild := 150

		treeroot.Insert(newNumber)
		treeroot.Insert(grandchild)

		gotValue := treeroot.GetRight().GetValue()
		grandChildValue := treeroot.GetRight().GetRight().GetValue()

		assertValue(t, newNumber, gotValue)

		assertValue(t, grandchild, grandChildValue)
	})
	t.Run("inserts an existing number would not create new node", func(t *testing.T) {
		treeroot := CreateBinaryTreeNode(100)
		newNumber := 120

		treeroot.Insert(newNumber)
		treeroot.Insert(newNumber)

		gotValue := treeroot.GetRight().GetValue()
		shouldBeNilRight := treeroot.GetRight().GetRight()
		shouldBeNilLeft := treeroot.GetRight().GetLeft()

		assertValue(t, newNumber, gotValue)

		if shouldBeNilLeft != nil || shouldBeNilRight != nil {
			t.Errorf("a duplicate node should not be added, but it was created %v, %v", shouldBeNilLeft, shouldBeNilRight)
		}
	})
	t.Run("searches a number to return path", func(t *testing.T) {
		treeroot := CreateBinaryTreeNode(100)
		child := 120
		grandchild := 150

		treeroot.Insert(child)
		treeroot.Insert(80)
		treeroot.Insert(115)
		treeroot.Insert(grandchild)
		treeroot.Insert(200)

		pathWanted := []*BinaryTree{treeroot, treeroot.GetRight(), treeroot.GetRight().GetRight()}

		pathGot, err := treeroot.LookFor(150)

		if err != nil {
			t.Fatalf("got an expected error: %v", err)
		}

		if !reflect.DeepEqual(pathWanted, pathGot) {
			t.Errorf("expected path to be %v, but got %v", pathWanted, pathGot)
		}
	})
	t.Run("searches an non-existent number to return error", func(t *testing.T) {
		treeroot := CreateBinaryTreeNode(99)
		child := 120
		grandchild := 150

		treeroot.Insert(child)
		treeroot.Insert(80)
		treeroot.Insert(115)
		treeroot.Insert(grandchild)
		treeroot.Insert(200)

		_, err := treeroot.LookFor(2)

		if err == nil {
			t.Fatalf("expected an error, but got none")
		}
	})
	t.Run("InOrderTraversal returns a string of InOrder", func(t *testing.T) {
		treeroot := CreateBinaryTreeNode(100)

		treeroot.Insert(120)
		treeroot.Insert(80)
		treeroot.Insert(40)
		treeroot.Insert(90)
		treeroot.Insert(115)
		treeroot.Insert(150)
		treeroot.Insert(200)

		expectedTraversal := "40 80 90 100 115 120 150 200"
		got := InOrderTraversal(treeroot)

		if expectedTraversal != got {
			t.Errorf("expected traversal to be %s, but got %s", expectedTraversal, got)
		}
	})
	// t.Run("GetInOrderSuccessor returns the smallest value on right branch", func(t *testing.T) {
	// 	treeroot := CreateBinaryTreeNode(100)
	// 	child := 120
	// 	grandchild := 150

	// 	treeroot.Insert(child)
	// 	treeroot.Insert(80)
	// 	treeroot.Insert(115)
	// 	treeroot.Insert(grandchild)
	// 	treeroot.Insert(200)

	// 	inOrderSuccessor := treeroot.GetInOrderSuccessor()

	// 	if inOrderSuccessor != 80 {
	// 		t.Errorf("expected inOrderSuccessor to be %d, but got %d", 30, inOrderSuccessor.value)
	// 	}
	// })
	// t.Run("removes a node with no child", func(t *testing.T) {
	// 	treeroot := CreateBinaryTreeNode(99)
	// 	child := 120
	// 	grandchild := 150

	// 	treeroot.Insert(child)
	// 	treeroot.Insert(80)
	// 	treeroot.Insert(115)
	// 	treeroot.Insert(grandchild)
	// 	treeroot.Insert(200)

	// 	treeroot.Delete(200)

	// 	_, err := treeroot.LookFor(200)

	// 	if err == nil {
	// 		t.Fatalf("expected an error, but got none")
	// 	}
	// })
}

func assertValue(t testing.TB, expectedValue, actualValue int) {
	t.Helper()
	if expectedValue != actualValue {
		t.Errorf("expected value to be %d, but got %d", expectedValue, actualValue)
	}
}
