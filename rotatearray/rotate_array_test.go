package rotatearray

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "one elt",
			nums: []int{0},
			k:    4,
			want: []int{0},
		},
		{
			name: "two elt",
			nums: []int{0, 1},
			k:    1,
			want: []int{1, 0},
		},
		{
			name: "three elts",
			nums: []int{1, 2, 3},
			k:    2,
			want: []int{2, 3, 1},
		},
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "Example 2",
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
