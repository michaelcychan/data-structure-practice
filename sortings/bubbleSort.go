package MySort

import "fmt"

func BubbleSort(inputSlice []int) []int {
	var counter int = 0
	for j := 0; j < len(inputSlice); j++ {
		for i := 0; i < len(inputSlice)-j-1; i++ {
			if inputSlice[i] > inputSlice[i+1] {
				swap(i, i+1, inputSlice)
				counter += 1
			}
		}
	}
	fmt.Printf("swapping done: %d \n", counter)
	return inputSlice
}

func swap(ind1, ind2 int, inputSlice []int) []int {
	temp := inputSlice[ind1]
	inputSlice[ind1] = inputSlice[ind2]
	inputSlice[ind2] = temp
	return inputSlice
}
