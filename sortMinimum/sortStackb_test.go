package sortMinimum

import "testing"

var TestCases4 = []struct {
	name string
	num int
	tempStack []int
	operations []string
	res1 []int
	res2 []string
}{
	{"Test1", 4, []int{5,3,1}, []string{"sa", "pb"}, []int{5,4,3,1}, []string{"sa", "pb", "sb"}},
}

func TestSortTempStack(t *testing.T) {
	for _ , tc := range TestCases4 {
		t.Run(tc.name, func(t *testing.T) {
			res1, res2 := SortTempStack(tc.num, tc.tempStack, tc.operations)

			for i := 0 ; i < len(res1) ; i++ {
				if res1[i] != tc.res1[i] {
					t.Errorf("Expected %v but got %v", tc.res1, res1)
				}
			}
			for i := 0 ; i < len(res2) ; i++ {
				if res2[i] != tc.res2[i] {
					t.Errorf("Expected %v but got %v", tc.res2, res2)
				}
			}
		})
	}
}
