package sortMinimum

// FindMax locates the maximum value and its index in the array.
func FindMax(arr []int) (maxVal, maxIndex int) {
	for i, val := range arr {
		if val > maxVal {
			maxVal = val
			maxIndex = i
		}
	}
	return maxVal, maxIndex
}
