package sortMinimum

// PushSwap handles the sorting process by pushing elements to a temporary stack.
func PushSwap(arr []int) ([]string, []int, []int) {
	var resl []string
	operations := []string{}
	tempStack := []int{}
	maxValue, maxIndex := FindMax(arr)
	count := 0

	for !IsSorted(arr) {
		count++
		if count > 5 {
			break
		}
		if len(arr) == 3 {
			var ops []string
			arr, ops = HandleThree(arr)
			operations = append(operations, ops...)
			break
		}

		if len(arr) > 1 && arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
			operations = append(operations, "sa")
			continue
		}

		if len(arr) > 1 && arr[0] > arr[len(arr)-1] {
			arr = RotateRight(arr)
			operations = append(operations, "rra")
			continue
		}

		check := maxIndex - 2
		if check >= 0 {
			if arr[check] != maxValue {
				maxValue, maxIndex = FindMax(arr)
			}
		}
		if maxIndex == 0 && len(arr) > 1 {
			arr = RotateLeft(arr)
			operations = append(operations, "ra")
			continue
		}

		arr, tempStack, resl = PushToB(arr, tempStack)
		operations = append(operations, "pb")
		operations = append(operations, resl...)
	}
	return operations, arr, tempStack
}
