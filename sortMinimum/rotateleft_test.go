package sortMinimum

import "testing"

var TestCases8 = []struct {
	name string
	arr  []int
	res  []int
}{
	{"Test1", []int{5, 1, 2, 3, 4}, []int{1, 2, 3, 4, 5}},
}

func TestRotateLeft(t *testing.T) {
	for _, tc := range TestCases8 {
		t.Run(tc.name, func(t *testing.T) {
			res := RotateLeft(tc.arr)

			for i := 0; i < len(res); i++ {
				if res[i] != tc.res[i] {
					t.Errorf("Expected %v but got %v", tc.res, res)
				}
			}
		})
	}
}
