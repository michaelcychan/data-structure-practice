package Queue

import "testing"

func TestQueue(t *testing.T) {
	t.Run("initialises", func(t *testing.T) {
		queue := MyQueue{cap: 5}
		_, err := queue.Peek()
		if err == nil {
			t.Error("Expected an error, but got none")
		}
	})
	t.Run("enqueue one data", func(t *testing.T) {
		queue := MyQueue{cap: 5}

		queue.Enqueue(3)
		expectedTop := 3

		top, err := queue.Peek()
		if err != nil {
			t.Errorf("expected no error, but got one: %v", err)
		}
		if top != expectedTop {
			t.Errorf("expected %d, but got %d", expectedTop, top)
		}
	})
}
