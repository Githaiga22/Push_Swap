package sortMinimum

// HandleThree sorts an array of three elements using minimal operations.
func HandleThree(input []int) ([]int, []string) {
	res := []string{}
	a, b, c := input[0], input[1], input[2]

	switch {
	case a > b && b < c && c > a:
		input[0], input[1] = input[1], input[0]
		res = append(res, "sa")
	case a > b && b > c:
		input[0], input[1] = input[1], input[0]
		res = append(res, "sa")
		input = RotateRight(input)
		res = append(res, "rra")
	case a > b && b < c && c < a:
		input = RotateLeft(input)
		res = append(res, "ra")
	case a < b && b > c && c > a:
		input[0], input[1] = input[1], input[0]
		res = append(res, "sa")
		input = RotateLeft(input)
		res = append(res, "ra")
	case a < b && b > c && c < a:
		input = RotateRight(input)
		res = append(res, "rra")
	}
	return input, res
}
