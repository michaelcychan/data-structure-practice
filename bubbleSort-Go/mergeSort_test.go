package MySort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	t.Run("returns a sorted slice (2 items)", func(t *testing.T) {
		original := []int{4, 2}
		expected := []int{2, 4}

		actual := mergeSort(original)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})
	t.Run("returns a sorted slice (10 items)", func(t *testing.T) {
		original := []int{6, 4, 2, 20, 16, 14, 18, 10, 12, 8}
		expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

		actual := mergeSort(original)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})
}

func TestMerge(t *testing.T) {
	t.Run("returns a combined slice from two single item slices", func(t *testing.T) {
		slice1 := []int{2}
		slice2 := []int{4}

		expected := []int{2, 4}
		got := merge(slice1, slice2)
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v, but got %v", expected, got)
		}
	})
	t.Run("returns a combined slice", func(t *testing.T) {
		slice1 := []int{2, 5, 9}
		slice2 := []int{4, 7, 10, 100}

		expected := []int{2, 4, 5, 7, 9, 10, 100}
		got := merge(slice1, slice2)
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v, but got %v", expected, got)
		}
	})
}
