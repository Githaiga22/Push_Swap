package sortMinimum

import "testing"

var TestCases7 = []struct {
	name  string
	input []int
	res1  []int
	res2  []string
}{
	{"Test1", []int{2, 3, 1}, []int{1, 2, 3}, []string{"rra"}},
}

func TestHandleThree(t *testing.T) {
	for _, tc := range TestCases7 {
		t.Run(tc.name, func(t *testing.T) {
			res1, res2 := HandleThree(tc.input)
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
