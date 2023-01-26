package BinaryTree

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
	} else if value < bt.value {
		if bt.left == nil {
			newNode := CreateBinaryTreeNode(value)
			bt.left = newNode
		} else {
			bt.left.Insert(value)
		}
	} else {
		fmt.Println(fmt.Errorf("%d already existed", value))
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

func InOrderTraversal(bt *BinaryTree) string {
	if bt == nil {
		return ""
	}
	traversalStack := []*BinaryTree{}
	note := []string{}
	curr := bt
	for {
		if curr != nil {
			traversalStack = append(traversalStack, curr)
			curr = curr.left
		} else if len(traversalStack) > 0 {
			popped := traversalStack[len(traversalStack)-1]
			traversalStack = traversalStack[:len(traversalStack)-1]
			poppedValue := popped.value
			note = append(note, strconv.Itoa(poppedValue))
			curr = popped.right
		} else {
			break
		}
	}

	return strings.Join(note, " ")

}

// func (bt *BinaryTree) GetInOrderSuccessor() *BinaryTree {
// 	if bt == nil {
// 		return nil
// 	}
// 	return nil
// }

// func (bt *BinaryTree) Delete(target int) {
// 	var parent, curr *BinaryTree
// 	if bt == nil {
// 		return
// 	}
// 	if target > bt.value {
// 		bt.right.Delete(target)
// 	} else if target < bt.value {
// 		bt.left.Delete(target)
// 	}
// }
