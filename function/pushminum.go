package push

import "fmt"

func PushMinimum(input []int) []string {
	var operations []string
	arr := append([]int{}, input...)
	var tempStack []int
	var maxVal, maxIndex, lastMax int

	for !IsSorted(arr) {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
			operations = append(operations, "sa")
			continue
		}

		for i, val := range arr {
			if val > maxVal && val != lastMax {
				maxVal = val
				maxIndex = i
			}
		}

		if maxIndex == len(arr)-1 {
			lastMax = maxVal
			maxVal = 0
			continue
		}

		for i := 0; i < maxIndex; i++ {
			if i < len(arr) {
				operations = append(operations, "pb")
				tempStack, operations = SortTempStack(arr[i], tempStack, operations)
			}
		}
		arr = arr[maxIndex:]
	}

	for _, val := range tempStack {
		arr = append([]int{val}, arr...)
		operations = append(operations, "pa")
	}

	fmt.Println(operations)
	fmt.Println(arr)
	fmt.Println(tempStack)

	return operations
}

func SortTempStack(num int, tempStack []int, operations []string) ([]int, []string) {
	if len(tempStack) == 0 {
		tempStack = append(tempStack, num)
		return tempStack, operations
	}

	if num > tempStack[0] {
		tempStack = append([]int{num}, tempStack...)
	} else if num < tempStack[len(tempStack)-1] {
		tempStack = append(tempStack, num)
		operations = append(operations, "rb")
	} else {
		tempStack = append(tempStack, num)
		tempStack[len(tempStack)-2], tempStack[len(tempStack)-1] = tempStack[len(tempStack)-1], tempStack[len(tempStack)-2]
		operations = append(operations, "sb")
	}

	return tempStack, operations
}
