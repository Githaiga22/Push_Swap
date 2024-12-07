package sortMinimum

import "testing"

var TestCases9 = []struct {
	name string
	arr  []int
	res  []int
}{
	{"Test1", []int{2, 3, 4, 5, 1}, []int{1, 2, 3, 4, 5}},
}

func TestRotateRight(t *testing.T) {
	for _, tc := range TestCases9 {
		t.Run(tc.name, func(t *testing.T) {
			res := RotateRight(tc.arr)

			for i := 0; i < len(res); i++ {
				if res[i] != tc.res[i] {
					t.Errorf("Expected %v but got %v", tc.res, res)
				}
			}
		})
	}
}
