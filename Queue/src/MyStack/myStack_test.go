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
	t.Run("adds one data", func(t *testing.T) {
		stack := MyStack{cap: 5}
		expectedTop := 6
		errAdd := stack.Add(expectedTop)

		if errAdd != nil {
			t.Fatalf("expected no error, but got one: %v", errAdd)
		}

		topValue, err := stack.Peek()
		if err != nil {
			t.Fatalf("Expected no error, but got one: %v", err)
		}

		if topValue != expectedTop {
			t.Errorf("expected %d, but got %d", expectedTop, topValue)
		}
	})
	t.Run("adds two data and Peek show the top", func(t *testing.T) {
		stack := MyStack{cap: 5}

		stack.Add(20)
		expectedTop := 6
		errAdd := stack.Add(expectedTop)

		if errAdd != nil {
			t.Fatalf("expected no error, but got one: %v", errAdd)
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

		stack.Add(54)
		stack.Add(20)
		expectedTop := 6
		stack.Add(expectedTop)

		errAdd := stack.Add(2)

		if errAdd == nil {
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
	t.Run("removeTop from empty stack warns of underflow", func(t *testing.T) {
		stack := MyStack{cap: 5}

		_, errUnderflow := stack.RemoveTop()

		if errUnderflow == nil {
			t.Fatalf("Expected an error, but got none")
		}
	})
	t.Run("removeTop from a stack with one data", func(t *testing.T) {
		stack := MyStack{cap: 5}
		expectedTop := 6
		stack.Add(expectedTop)

		oldTop, err := stack.RemoveTop()

		if err != nil {
			t.Fatalf("expected no error, but got one %v", err)
		}

		if oldTop != expectedTop {
			t.Errorf("expected %d, but got %d", expectedTop, oldTop)
		}
	})
	t.Run("removeTop from a stack with two data", func(t *testing.T) {
		stack := MyStack{cap: 5}
		expectedTop := 6
		bigBase := 100
		stack.Add(bigBase)
		stack.Add(expectedTop)

		oldTop, err := stack.RemoveTop()

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
		stack.Add(240)

		stack.Add(73)

		stack.RemoveTop()
		stack.RemoveTop()
		stack.Add(bigBase)
		stack.Add(90)
		stack.RemoveTop()
		stack.Add(expectedTop)

		oldTop, err := stack.RemoveTop()

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
