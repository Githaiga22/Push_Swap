package sortMinimum

// SortTempStack organizes the temp stack based on the current number.
func SortTempStack(num int, tempStack []int, operations []string) ([]int, []string) {
	if len(tempStack) == 0 || num > tempStack[0] {
		tempStack = append([]int{num}, tempStack...)
	} else {
		tempStack = append([]int{num}, tempStack...)
		if len(tempStack) > 1 && tempStack[1] > tempStack[0] {
			tempStack[0], tempStack[1] = tempStack[1], tempStack[0]
			operations = append(operations, "sb")
		}
		if len(tempStack) > 1 && tempStack[0] < tempStack[len(tempStack)-1] {
			tempStack = RotateLeft(tempStack)
			operations = append(operations, "rb")
		}
		if len(tempStack) > 1 && tempStack[len(tempStack)-1] > tempStack[0] {
			tempStack = RotateRight(tempStack)
			operations = append(operations, "rrb")
		}
	}
	return tempStack, operations
}
