package MySort

func QuickSort(inputSlice []int) []int {
	if len(inputSlice) <= 1 {
		return inputSlice
	}
	var pivot int = inputSlice[0]
	lt := []int{}
	gt := []int{}
	for i := 1; i < len(inputSlice); i++ {
		if inputSlice[i] > pivot {
			gt = append(gt, inputSlice[i])
		} else {
			lt = append(lt, inputSlice[i])
		}
	}
	result := append(QuickSort(lt), pivot)
	result = append(result, QuickSort(gt)...)

	return result
}
