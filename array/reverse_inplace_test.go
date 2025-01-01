package array

import (
	"slices"
	"testing"
)

func TestReverseInPlace(t *testing.T) {
	tests := []struct {
		list    []int
		start   int
		end     int
		reverse []int
	}{
		{[]int{}, 0, 0, []int{}},
		{[]int{1, 2, 3, 4, 5}, 1, 4, []int{1, 5, 4, 3, 2}},
		{[]int{1, 2, 3, 4, 5, 6}, 0, 3, []int{4, 3, 2, 1, 5, 6}},
	}

	for i, test := range tests {
		ReverseInPlace(test.list, test.start, test.end)
		if !slices.Equal(test.list, test.reverse) {
			t.Fatalf("failed test case #%d. want %#v got %#v", i, test.reverse, test.list)
		}
	}

}
