package sortMinimum

import "fmt"

// PushMinimum finds the sequence of operations to sort the array.
func PushMinimum(input []int) []string {
	operations := []string{}
	arr := append([]int{}, input...)

	for !IsSorted(arr) {
		ops, updatedArr, tempStack := PushSwap(arr)
		operations = append(operations, ops...)
		resl, finalArr := ProcessTempStack(updatedArr, tempStack)
		operations = append(operations, resl...)
		arr = finalArr
	}

	operations = SortOperations(operations)
	fmt.Println("Final Array:", arr)
	fmt.Println("Original Input:", input)
	fmt.Println("Operations:", operations)
	return operations
}
