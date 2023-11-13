package main

import "testing"

var tests = []struct {
	name string
	nums []int
	val  int
	want int
}{
	{
		name: "Example 1",
		nums: []int{3, 2, 2, 3},
		val:  3,
		want: 2,
	},
	{
		name: "Example 2",
		nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
		val:  2,
		want: 5,
	},
	{
		name: "Empty nums",
		nums: []int{},
		val:  0,
		want: 0,
	},
	{
		name: "nothing to delete",
		nums: []int{2},
		val:  3,
		want: 1,
	},
}

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.nums, tt.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_removeElementHiden(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElementHidden(tt.nums, tt.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeElement(tests[0].nums, tests[0].val)
	}
}

func BenchmarkRemoveElementHidden(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeElement(tests[0].nums, tests[0].val)
	}
}
