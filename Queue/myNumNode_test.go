package myNode

import "testing"

func TestMyNodes(t *testing.T) {
	t.Run("initialises", func(t *testing.T) {
		testNode := MyNode{value: 10}

		actualValue := testNode.GetValue()
		actuaNextNode := testNode.GetNextNode()

		expectedValue := 10
		var expectedNextNode *MyNode = nil

		if actualValue != expectedValue {
			t.Errorf("Value: expected %d, but got %d", expectedValue, actualValue)
		}

		if actuaNextNode != expectedNextNode {
			t.Errorf("Next Node: expected %v, but got %v", &expectedNextNode, &actuaNextNode)
		}
	})
}
