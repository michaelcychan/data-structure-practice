package MyStack

import "testing"

func TestMyStack(t *testing.T) {
	t.Run("initialises with an empty stack", func(t *testing.T) {
		stack := MyStack{}

		_, err := stack.Peek()
		if err == nil {
			t.Fatalf("Expected an error, but got none")
		}
	})
}
