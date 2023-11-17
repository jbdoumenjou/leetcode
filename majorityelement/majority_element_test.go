package majorityelement

import "testing"

func Test_majorityElement(t *testing.T) {
	testcases := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 0",
			nums: []int{1},
			want: 1,
		},
		{
			name: "Example 1",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{3, 3, 4},
			want: 3,
		},
		{
			name: "Example 3",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			name: "Example 4",
			nums: []int{-1, 1, 1, 1, 2, 1},
			want: 1,
		},
	}

	for _, tt := range testcases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := majorityElement(tt.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
