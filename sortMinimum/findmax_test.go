package sortMinimum

import "testing"

var TestCases6 = []struct {
	name     string
	arr      []int
	maxVal   int
	maxIndex int
}{
	{"Test1", []int{2, 3, 5, 6, 9}, 9, 4},
	{"Test2", []int{5, 9, 2, 3, 1}, 9, 1},
}

func TestFindMax(t *testing.T) {
	for _, tc := range TestCases6 {
		t.Run(tc.name, func(t *testing.T) {
			maxVal, maxIndex := FindMax(tc.arr)

			if maxIndex != tc.maxIndex || maxVal != tc.maxVal {
				t.Errorf("Expected the value %v and index %v but got value %v and index %v", tc.maxVal, tc.maxIndex, maxVal, maxIndex)
			}
		})
	}
}
