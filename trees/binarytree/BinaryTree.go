package BinaryTree

import (
	"errors"
)

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func CreateBinaryTreeNode(value int) *BinaryTree {
	return &BinaryTree{value: value}
}

func (bt *BinaryTree) GetValue() int {
	return bt.value
}

func (bt *BinaryTree) GetLeft() *BinaryTree {
	return bt.left
}

func (bt *BinaryTree) GetRight() *BinaryTree {
	return bt.right
}

func (bt *BinaryTree) Insert(value int) {
	if value > bt.value {
		if bt.right == nil {
			newNode := CreateBinaryTreeNode(value)
			bt.right = newNode
		} else {
			bt.right.Insert(value)
		}
	} else {
		if bt.left == nil {
			newNode := CreateBinaryTreeNode(value)
			bt.left = newNode
		} else {
			bt.right.Insert(value)
		}
	}
}

func (bt *BinaryTree) LookFor(target int) ([]*BinaryTree, error) {
	path := []*BinaryTree{bt}

	if target > bt.value {
		if bt.right != nil {
			returnedPath, err := bt.right.LookFor(target)
			if err != nil {
				return nil, err
			} else {
				path = append(path, returnedPath...)
			}
		} else {
			return nil, errors.New("not found")
		}

	} else if target < bt.value {
		if bt.left != nil {
			returnedPath, err := bt.left.LookFor(target)
			if err != nil {
				return nil, err
			} else {
				path = append(path, returnedPath...)
			}
		} else {
			return nil, errors.New("not found")
		}
	} else {
		return path, nil
	}
	return path, nil
}
