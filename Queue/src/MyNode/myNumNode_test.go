package MyNode

import (
	"testing"
)

func TestMyNodes(t *testing.T) {
	t.Run("initialises", func(t *testing.T) {
		testNode := MyNode{value: 10}

		actualValue := testNode.GetValue()
		actuaNextNode := testNode.GetNext()

		expectedValue := 10
		var expectedNextNode *MyNode = nil

		if actualValue != expectedValue {
			t.Errorf("Value: expected %d, but got %d", expectedValue, actualValue)
		}

		if actuaNextNode != expectedNextNode {
			t.Errorf("Next Node: expected %v, but got %v", &expectedNextNode, &actuaNextNode)
		}
	})
	t.Run("initialises with no input", func(t *testing.T) {
		testNode := MyNode{}

		actualValue := testNode.GetValue()
		actuaNextNode := testNode.GetNext()

		expectedValue := 0
		var expectedNextNode *MyNode = nil

		if actualValue != expectedValue {
			t.Errorf("Value: expected %d, but got %d", expectedValue, actualValue)
		}

		if actuaNextNode != expectedNextNode {
			t.Errorf("Next Node: expected %v, but got %v", &expectedNextNode, &actuaNextNode)
		}
	})
	t.Run("linking two nodes", func(t *testing.T) {
		node1 := MyNode{value: 3}
		node2 := MyNode{value: 13}

		node1.SetNext(node2)

		expectedNextNodeValue := 13
		actualNextNodeVallue := node1.GetNext().GetValue()

		if expectedNextNodeValue != actualNextNodeVallue {
			t.Errorf("expected next node should be %d, but got %d", expectedNextNodeValue, actualNextNodeVallue)
		}
	})

	t.Run("linking two nodes then delink", func(t *testing.T) {
		node1 := MyNode{value: 3}
		node2 := MyNode{value: 13}

		node1.SetNext(node2)

		expectedNextNodeValue := 13
		actualNextNodeVallue := node1.GetNext().GetValue()

		if expectedNextNodeValue != actualNextNodeVallue {
			t.Errorf("expected next node should be %d, but got %d", expectedNextNodeValue, actualNextNodeVallue)
		}

		node1.RemoveNext()

		var expectedEmptyNext *MyNode = nil
		actualNodeAfterRemoval := node1.GetNext()

		if expectedEmptyNext != actualNodeAfterRemoval {
			t.Errorf("expected next node to be %v, but got %v", expectedEmptyNext, actualNodeAfterRemoval)
		}

	})
}
