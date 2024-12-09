package sortMinimum

import "testing"

var TestCases11 = []struct {
	name string
	arr  []int
	res1 []string
	res2 []int
	res3 []int
}{
	{"Test1", []int{2, 1, 3, 6, 5, 8}, []string{"sa", "pb", "pb", "pb", "sa"}, []int{5, 6, 8}, []int{3, 2, 1}},
}

func TestPushSwap(t *testing.T) {
	for _, tc := range TestCases11 {
		t.Run(tc.name, func(t *testing.T) {
			res1, res2, res3 := PushSwap(tc.arr)
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
			for i := 0; i < len(res3); i++ {
				if res3[i] != tc.res3[i] {
					t.Errorf("Expected %v but got %v", tc.res3, res3)
				}
			}
		})
	}
}
