package sortMinimum

// ProcessTempStack reintegrates the temporary stack into the array.
func ProcessTempStack(arr []int, tempStack []int) ([]string, []int) {
	operations := []string{}
	for len(tempStack) > 0 {
		val := tempStack[0]
		tempStack = tempStack[1:]
		arr = append([]int{val}, arr...)
		operations = append(operations, "pa")
		if len(arr) > 1 && arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
			operations = append(operations, "sa")
		}
	}
	return operations, arr
}
