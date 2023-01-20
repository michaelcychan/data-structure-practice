package MySort

func MergeSort(inputSlice []int) []int {
	if len(inputSlice) <= 1 {
		return inputSlice
	} else {
		middleIndex := len(inputSlice) / 2
		leftSlice := inputSlice[:middleIndex]
		rightSlice := inputSlice[middleIndex:]
		return merge(MergeSort(leftSlice), MergeSort(rightSlice))
	}
}

func merge(slice1, slice2 []int) []int {
	result := []int{}
	for len(slice1) > 0 && len(slice2) > 0 {
		if slice1[0] < slice2[0] {
			result = append(result, slice1[0])
			slice1 = slice1[1:]
		} else {
			result = append(result, slice2[0])
			slice2 = slice2[1:]
		}
	}
	if len(slice1) > 0 {
		result = append(result, slice1...)
	}
	if len(slice2) > 0 {
		result = append(result, slice2...)
	}
	return result
}
