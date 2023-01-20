package MySort

import (
	"fmt"
	"math/rand"
	"time"
)

// decoupling QuickSort and qsort follows the pricinple of encapsulation
// users are not required to input their own start and end fields,
// simplying the usage also reducing errors
func QuickSort(inputSlice []int) []int {
	// start and end are used as index
	start := 0
	end := len(inputSlice) - 1

	qsort(&inputSlice, start, end)
	return inputSlice
}

// takes pointer => pass by reference instead of value
// this reduces the need to create separate slices,
// so that the space complexity remains O(N)
func qsort(inputSlice *[]int, start, end int) {
	fmt.Printf("start: %d, end: %d \n", start, end)
	if start < end {

		// pivot will be the comparison pivot for partitioning
		// proper seeding is needed to have a different rand.Intn() each time
		// pivotIdx has to be between start and end
		rand.Seed(time.Now().UTC().UnixNano())
		pivotIdx := rand.Intn((end-start)+1) + start
		var pivot int = (*inputSlice)[pivotIdx]

		// swapping pivot with the end element
		(*inputSlice)[pivotIdx], (*inputSlice)[end] = (*inputSlice)[end], (*inputSlice)[pivotIdx]

		// creating less-than-pivot-index
		ltpIdx := start

		for i := start; i <= end; i++ {
			// anything smaller than pivot will be swapped to front
			if (*inputSlice)[i] < pivot {
				(*inputSlice)[i], (*inputSlice)[ltpIdx] = (*inputSlice)[ltpIdx], (*inputSlice)[i]

				// after swapping, ltpIdx will move for the next potential swap
				ltpIdx++
			}
		}

		// swap pivot with marker position
		// end result: elements before marker (pivot) are smaller than pivot
		(*inputSlice)[ltpIdx], (*inputSlice)[end] = (*inputSlice)[end], (*inputSlice)[ltpIdx]

		// recursively sorting the two partitions before and after the pivot
		qsort(inputSlice, start, ltpIdx-1)
		qsort(inputSlice, ltpIdx+1, end)
	}
}
