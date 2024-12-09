package sortMinimum

// RotateRight performs a right rotation on the array.
func RotateRight(arr []int) []int {
	return append([]int{arr[len(arr)-1]}, arr[:len(arr)-1]...)
}
