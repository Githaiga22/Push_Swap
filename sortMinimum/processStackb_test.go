package sortMinimum

import "testing"

var TestCases3 = []struct {
	name      string
	arr       []int
	tempStack []int
	operation []string
	res       []int
}{
	{"Test1", []int{7, 8}, []int{5, 1}, []string{"pa", "pa"}, []int{1, 5, 7, 8}},
	{"Test2", []int{4, 7, 8}, []int{5, 1}, []string{"pa", "sa", "pa"}, []int{1, 4, 5, 7, 8}},
}

func TestProcessTempStack(t *testing.T) {
	for _, tc := range TestCases3 {
		t.Run(tc.name, func(t *testing.T) {
			ops, arr := ProcessTempStack(tc.arr, tc.tempStack)

			for i := 0; i < len(arr); i++ {
				if arr[i] != tc.res[i] {
					t.Errorf("Expected %v but got %v", tc.res[i], arr[i])
				}
			}

			for i := 0; i < len(ops); i++ {
				if ops[i] != tc.operation[i] {
					t.Errorf("Expected %v but got %v", tc.operation[i], ops[i])
				}
			}
		})
	}
}
