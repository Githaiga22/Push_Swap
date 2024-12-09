package sortMinimum

// RotateLeft performs a left rotation on the array.
func RotateLeft(arr []int) []int {
	return append(arr[1:], arr[0])
}
