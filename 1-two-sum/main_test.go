package main

import (
	"reflect"
	"testing"
)

// TestTwoSum
func TestTwoSum(t *testing.T) {
	type test struct {
		nums   []int
		target int
		want   []int
	}

	tests := []test{
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}

	for _, tc := range tests {
		got := twoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
