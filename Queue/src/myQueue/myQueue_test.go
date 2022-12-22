package Queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("initialises", func(t *testing.T) {
		expectedCap := 5

		queue := MyQueue{cap: expectedCap}

		actualCap := queue.GetCap()

		if expectedCap != actualCap {
			t.Errorf("Expected cap to be %d, but got %d", expectedCap, actualCap)
		}

		expectedSize := 0
		actualSize := queue.CurrentSize()

		if expectedSize != actualSize {
			t.Errorf("Expected current size to be %d, but got %d", expectedSize, actualSize)
		}

		_, err := queue.Peek()
		if err == nil {
			t.Error("Expected an error, but got none")
		}
	})
	t.Run("enqueue one data", func(t *testing.T) {
		queue := MyQueue{cap: 5}

		expectedTop := 3
		errEnqueue := queue.Enqueue(expectedTop)

		assertNoError(t, errEnqueue)

		top, err := queue.Peek()

		assertNoError(t, err)
		if top != expectedTop {
			t.Errorf("expected %d, but got %d", expectedTop, top)
		}

		expectedSize := 1
		actualSize := queue.CurrentSize()
		if expectedSize != actualSize {
			t.Errorf("Expected current size at %d, but got %d", expectedSize, actualSize)
		}
	})
	t.Run("enqueue two data", func(t *testing.T) {
		queue := MyQueue{cap: 5}

		expectedTop := 3
		queue.Enqueue(expectedTop)
		errEnqueue := queue.Enqueue(17)

		assertNoError(t, errEnqueue)

		top, err := queue.Peek()
		assertNoError(t, err)

		if top != expectedTop {
			t.Errorf("expected %d, but got %d", expectedTop, top)
		}

		expectedSize := 2
		actualSize := queue.CurrentSize()
		if expectedSize != actualSize {
			t.Errorf("Expected current size at %d, but got %d", expectedSize, actualSize)
		}
	})
	t.Run("enqueue one data and dequeue", func(t *testing.T) {
		queue := MyQueue{cap: 5}

		expectedTop := 3
		queue.Enqueue(expectedTop)

		queue.Dequeue()

		expectedSize := 0
		actualSize := queue.CurrentSize()
		if expectedSize != actualSize {
			t.Errorf("Expected current size at %d, but got %d", expectedSize, actualSize)
		}

		_, err := queue.Peek()

		assertError(t, err)
	})
	t.Run("enqueue two data and dequeue one", func(t *testing.T) {
		queue := MyQueue{cap: 5}

		expectedTop := 3
		expectedBottom := 17

		queue.Enqueue(expectedTop)
		queue.Enqueue(expectedBottom)

		actualRemoved, errDequeue := queue.Dequeue()

		assertNoError(t, errDequeue)

		if actualRemoved != expectedTop {
			t.Errorf("expected %d to be removed, but got %d", expectedTop, actualRemoved)
		}

		expectedSize := 1
		actualSize := queue.CurrentSize()
		if expectedSize != actualSize {
			t.Errorf("Expected current size at %d, but got %d", expectedSize, actualSize)
		}

		finalTop, errPeak := queue.Peek()

		assertNoError(t, errPeak)

		if finalTop != expectedBottom {
			t.Errorf("expected %d to remain in queue, but got %d", expectedBottom, finalTop)
		}
	})
	t.Run("Enqueue when cap is reached", func(t *testing.T) {
		queue := MyQueue{cap: 1}
		errFirstEnqueue := queue.Enqueue(1)
		assertNoError(t, errFirstEnqueue)
		errCapReached := queue.Enqueue(2)
		assertError(t, errCapReached)
	})
	t.Run("Dequeue an empty list", func(t *testing.T) {
		queue := MyQueue{cap: 1}
		value := 1
		queue.Enqueue(value)
		valDequeue, errFirstDequeue := queue.Dequeue()

		assertNoError(t, errFirstDequeue)
		if valDequeue != value {
			t.Errorf("expected %d, but got %d", value, valDequeue)
		}
		_, errUnderflow := queue.Dequeue()
		assertError(t, errUnderflow)
	})
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error, but got one: %v", err)
	}
}

func assertError(t testing.TB, actualError error) {
	t.Helper()
	if actualError == nil {
		t.Fatalf("expected an error, but got none")
	}
}
