package push

import (
	"testing"
)

var TestCases1 = []struct {
	name      string
	input     string
	expected1 []int
	expected2 string
}{
	{"Test1", "2 3 4 5 6 7", []int{2, 3, 4, 5, 6, 7}, ""},
	{"Test2", "1 8 9 4 5 6 7", []int{1, 8, 9, 4, 5, 6, 7}, ""},
	{"Test3", "1 ok 2 3 4", nil, "not a number"},
}

func TestToArray(t *testing.T) {
	for _, tc := range TestCases1 {
		t.Run(tc.name, func(t *testing.T) {
			expected1 := tc.expected1
			expected2 := tc.expected2
			got1, got2 := ToArray(tc.input)

			if got2 != expected2 {
				t.Errorf("Expected %v and %v but got %v and %v", expected1, expected2, got1, got2)
			}
		})
	}
}
