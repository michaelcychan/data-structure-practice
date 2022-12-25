package MyStack

import "testing"

func TestMyStack(t *testing.T) {
	t.Run("initialises with an empty stack", func(t *testing.T) {
		stack := MyStack{cap: 5}

		_, err := stack.Peek()
		if err == nil {
			t.Fatalf("Expected an error, but got none")
		}
	})
	t.Run("Pushes one data", func(t *testing.T) {
		stack := MyStack{cap: 5}
		expectedTop := 6
		errPush := stack.Push(expectedTop)

		if errPush != nil {
			t.Fatalf("expected no error, but got one: %v", errPush)
		}

		topValue, err := stack.Peek()
		if err != nil {
			t.Fatalf("Expected no error, but got one: %v", err)
		}

		if topValue != expectedTop {
			t.Errorf("expected %d, but got %d", expectedTop, topValue)
		}
	})
	t.Run("Pushes two data and Peek show the top", func(t *testing.T) {
		stack := MyStack{cap: 5}

		stack.Push(20)
		expectedTop := 6
		errPush := stack.Push(expectedTop)

		if errPush != nil {
			t.Fatalf("expected no error, but got one: %v", errPush)
		}

		topValue, err := stack.Peek()
		if err != nil {
			t.Fatalf("Expected no error, but got one: %v", err)
		}

		if topValue != expectedTop {
			t.Errorf("expected %d, but got %d", expectedTop, topValue)
		}
	})
	t.Run("throws error if stack overflow", func(t *testing.T) {
		stack := MyStack{cap: 3}

		stack.Push(54)
		stack.Push(20)
		expectedTop := 6
		stack.Push(expectedTop)

		errPush := stack.Push(2)

		if errPush == nil {
			t.Fatalf("expected an error, but got none")
		}

		topValue, err := stack.Peek()
		if err != nil {
			t.Fatalf("Expected no error, but got one: %v", err)
		}

		if topValue != expectedTop {
			t.Errorf("expected %d, but got %d", expectedTop, topValue)
		}
	})
	t.Run("Pop from empty stack warns of underflow", func(t *testing.T) {
		stack := MyStack{cap: 5}

		_, errUnderflow := stack.Pop()

		if errUnderflow == nil {
			t.Fatalf("Expected an error, but got none")
		}
	})
	t.Run("Pop from a stack with one data", func(t *testing.T) {
		stack := MyStack{cap: 5}
		expectedTop := 6
		stack.Push(expectedTop)

		oldTop, err := stack.Pop()

		if err != nil {
			t.Fatalf("expected no error, but got one %v", err)
		}

		if oldTop != expectedTop {
			t.Errorf("expected %d, but got %d", expectedTop, oldTop)
		}
	})
	t.Run("Pop from a stack with two data", func(t *testing.T) {
		stack := MyStack{cap: 5}
		expectedTop := 6
		bigBase := 100
		stack.Push(bigBase)
		stack.Push(expectedTop)

		oldTop, err := stack.Pop()

		if err != nil {
			t.Fatalf("expected no error, but got one %v", err)
		}

		if oldTop != expectedTop {
			t.Errorf("expected %d, but got %d", expectedTop, oldTop)
		}

		newTop, errPeek := stack.Peek()
		if errPeek != nil {
			t.Fatalf("expected no error, but got one: %v", errPeek)
		}

		if newTop != bigBase {
			t.Errorf("expected %d, but got %d", bigBase, newTop)
		}
	})
	t.Run("a random case", func(t *testing.T) {
		stack := MyStack{cap: 5}
		expectedTop := 6
		bigBase := 100
		stack.Push(240)

		stack.Push(73)

		stack.Pop()
		stack.Pop()
		stack.Push(bigBase)
		stack.Push(90)
		stack.Pop()
		stack.Push(expectedTop)

		oldTop, err := stack.Pop()

		if err != nil {
			t.Fatalf("expected no error, but got one %v", err)
		}

		if oldTop != expectedTop {
			t.Errorf("expected %d, but got %d", expectedTop, oldTop)
		}

		newTop, errPeek := stack.Peek()
		if errPeek != nil {
			t.Fatalf("expected no error, but got one: %v", errPeek)
		}

		if newTop != bigBase {
			t.Errorf("expected %d, but got %d", bigBase, newTop)
		}
	})
}
