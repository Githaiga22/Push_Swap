package sortMinimum

import "testing"

var TestCases = []struct {
	name     string
	input    []int
	expected bool
}{
	{"Test1", []int{1, 2, 3, 4, 5}, true},
	{"Test2", []int{5, 4, 1, 9, 2}, false},
	{"Test3", []int{8, 99, 100, 103, 600}, true},
}

func TestCheckStack(t *testing.T) {
	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := tc.expected
			got := IsSorted(tc.input)

			if expected != got {
				t.Errorf("Expected %v but got %v", expected, got)
			}
		})
	}
}
