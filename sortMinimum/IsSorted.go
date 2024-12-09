package sortMinimum

import "sort"

func IsSorted(input []int) bool {
	return sort.IntsAreSorted(input)
}
