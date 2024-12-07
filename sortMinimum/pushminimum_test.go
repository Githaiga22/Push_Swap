package sortMinimum

import "testing"

var TestCases2 = []struct {
	name     string
	input    []int
	expected int
}{
	{"Test1", []int{2, 1, 3, 6, 5, 8}, 9},
	{"Test2", []int{2, 4, 5, 1, 7, 8}, 12},
	{"Test3", []int{246, 600, 166, 902, 106}, 12},
	{"Test4", []int{190, 331, 548, 818, 436}, 12},
	{"Test5", []int{560, 913, 721, 87, 575}, 12},
	{"Test6", []int{924, 169, 992, 451, 685}, 12},
	{"Test7", []int{821, 739, 866, 160, 388}, 12},
	{"Test8", []int{0, 1, 2, 3, 4, 5}, 0},
	{"Test9", []int{21, 23, 84, 31, 30}, 12},
	{"Test10", []int{7, 39, 90, 79, 94}, 12},
	{"Test11", []int{85, 28, 75, 93, 79}, 12},
	{"Test12", []int{15, 69, 35, 90, 34}, 12},
	{"Test13", []int{61, 34, 57, 63, 23}, 12},
	{"Test15", []int{61, 66, 63, 14, 16}, 12},
	{"Test14", []int{9, 1, 6, 5, 8}, 12},
	{"Test16", []int{12, 86, 74, 57, 73}, 12},
	{"Test17", []int{90, 82, 24, 49, 74}, 12},
	{"Test18", []int{27, 35, 19, 30, 28}, 12},
	{"Test19", []int{3, 7, 12, 19, 25}, 12},
	{"Test20", []int{8, 15, 22, 31, 40}, 12},
	{"Test21", []int{4, 9, 14, 18, 23}, 12},
	{"Test22", []int{1, 6, 11, 17, 24}, 12},
	{"Test23", []int{2, 5, 10, 16, 21}, 12},

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
