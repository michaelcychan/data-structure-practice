package MySort

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	t.Run("swap two elements in a slice", func(t *testing.T) {
		original := []int{1, 3, 5, 7, 9}
		expected := []int{1, 9, 5, 7, 3}

		actual := swap(1, 4, original)
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected result %v, but got %v", expected, actual)
		}

	})
}

func TestBubbleSort(t *testing.T) {
	t.Run("returns a sorted slice (2 items)", func(t *testing.T) {
		original := []int{4, 2}
		expected := []int{2, 4}

		actual := bubbleSort(original)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v, but got %v", expected, actual)
		}

	})
	t.Run("returns a sorted slice (3 items)", func(t *testing.T) {
		original := []int{6, 4, 2}
		expected := []int{2, 4, 6}

		actual := bubbleSort(original)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})
	t.Run("returns a sorted slice (10 items)", func(t *testing.T) {
		original := []int{6, 4, 2, 20, 16, 14, 18, 10, 12, 8}
		expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

		actual := bubbleSort(original)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})
}
