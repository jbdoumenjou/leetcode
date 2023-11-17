package runningsum

import (
	"reflect"
	"testing"
)

func Test_runningSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "simple test",
			nums: []int{1, 2, 3, 4},
			want: []int{1, 3, 6, 10},
		},
		{
			name: "all the same",
			nums: []int{1, 1, 1, 1},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "mixed values",
			nums: []int{3, 1, 2, 10, 1},
			want: []int{3, 4, 6, 16, 17},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if !reflect.DeepEqual(runningSum(tt.nums), tt.want) {
				t.Errorf("runningSum() = %v, want %v", runningSum(tt.nums), tt.want)
			}

		})
	}
}
func Test_runningSum_clone(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "simple test",
			nums: []int{1, 2, 3, 4},
			want: []int{1, 3, 6, 10},
		},
		{
			name: "all the same",
			nums: []int{1, 1, 1, 1},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "mixed values",
			nums: []int{3, 1, 2, 10, 1},
			want: []int{3, 4, 6, 16, 17},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := runningSum_clone(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runningSum() = %v, want %v", got, tt.want)
			}

			if reflect.DeepEqual(tt.nums, got) {
				t.Errorf("runningSum() mutated input: %v", tt.nums)
			}
		})
	}
}
