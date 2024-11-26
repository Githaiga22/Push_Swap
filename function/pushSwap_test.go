package push

import "testing"

var TestCases2 = []struct {
	name     string
	input    []int
	expected int
}{
	{"Test1", []int{2, 1, 3, 6, 5, 8}, 9},
	{"Test2", []int{2,4,5,1,7,8}, 12},
}

func TestPushMinum(t *testing.T) {
	for _, tc := range TestCases2 {
		t.Run(tc.name, func(t *testing.T) {
			expected := tc.expected
			got := len(PushMinimum(tc.input))

			if got > expected {
				t.Errorf("Expected %v but got %v", expected, got)
			}
		})
	}
}
