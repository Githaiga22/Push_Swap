package sortMinimum

import "testing"

var TestCases10 = []struct {
	name  string
	input []string
	res   []string
}{
	{"Test1", []string{"sb","sa"}, []string{"ss"}},
}

func TestSortOperations(t *testing.T) {
	for _, tc := range TestCases10 {
		t.Run(tc.name, func(t *testing.T) {
			res := SortOperations(tc.input)
			for i := 0; i < len(res); i++ {
				if res[i] != tc.res[i] {
					t.Errorf("Expected %v but got %v", tc.res, res)
				}
			}
		})
	}
}
