package sortMinimum

func PushToB(arr []int, tempStack []int) ([]int, []int, []string) {
	var operations []string
	tempStack, operations = SortTempStack(arr[0], tempStack, operations)
	arr = arr[1:]
	return arr, tempStack, operations
}
