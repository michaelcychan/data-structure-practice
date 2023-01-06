package MySort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("returns a sorted slice (2 items)", func(t *testing.T) {
		original := []int{4, 2}
		expected := []int{2, 4}

		actual := QuickSort(original)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})
	t.Run("returns a sorted slice (3 items)", func(t *testing.T) {
		original := []int{6, 4, 2}
		expected := []int{2, 4, 6}

		actual := MergeSort(original)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})
	t.Run("returns a sorted slice (10 items)", func(t *testing.T) {
		original := []int{6, 4, 2, 20, 16, 14, 18, 10, 12, 8}
		expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

		actual := QuickSort(original)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})
}
