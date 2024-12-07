package sortMinimum

import "testing"

var TestCases5 = []struct {
	name      string
	arr       []int
	tempStack []int
	res1      []int
	res2      []int
}{
	{"Test1", []int{5, 4, 6, 1}, []int{3}, []int{4, 6, 1}, []int{5, 3}},
}

func TestPushToB(t *testing.T) {
	for _, tc := range TestCases5 {
		t.Run(tc.name, func(t *testing.T) {
			res1, res2, _ := PushToB(tc.arr, tc.tempStack)

			for i := 0; i < len(res1); i++ {
				if res1[i] != tc.res1[i] {
					t.Errorf("Expected %v but got %v", tc.res1, res1)
				}
			}

			for i := 0; i < len(res2); i++ {
				if res2[i] != tc.res2[i] {
					t.Errorf("Expected %v but got %v", tc.res2, res2)
				}
			}
		})
	}
}
