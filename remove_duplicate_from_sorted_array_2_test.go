package main

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates2(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		expectedLen   int
		expectedArray []int
	}{
		{
			name:          "Example 1",
			nums:          []int{1, 1, 2},
			expectedLen:   3,
			expectedArray: []int{1, 1, 2},
		},
		{
			name:          "Example 2",
			nums:          []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedLen:   9,
			expectedArray: []int{0, 0, 1, 1, 2, 2, 3, 3, 4},
		},
		{
			name:          "Example 3",
			nums:          []int{1, 1, 2, 3},
			expectedLen:   4,
			expectedArray: []int{1, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates2(tt.nums); got != tt.expectedLen {
				t.Errorf("removeDuplicates() = %v, expectedLen %v", got, tt.expectedLen)
			}
			if !reflect.DeepEqual(tt.nums[:tt.expectedLen], tt.expectedArray) {
				t.Errorf("got = %v, expectedArray %v", tt.nums[:tt.expectedLen], tt.expectedArray)
			}
		})
	}
}
