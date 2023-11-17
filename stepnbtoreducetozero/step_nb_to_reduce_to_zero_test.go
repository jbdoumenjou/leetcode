package stepnbtoreducetozero

import "testing"

var testCases = []struct {
	name string
	num  int
	want int
}{
	{
		name: "example 1",
		num:  14,
		want: 6,
	},
	{
		name: "example 2",
		num:  8,
		want: 4,
	},
	{
		name: "example 3",
		num:  123,
		want: 12,
	},
	{
		name: "zero",
		num:  0,
		want: 0,
	},
}

func Test_numberOfSteps(t *testing.T) {
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := numberOfSteps(tt.num); got != tt.want {
				t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberOfSteps_bit(t *testing.T) {
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := numberOfSteps_bit(tt.num); got != tt.want {
				t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_numberOfSteps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberOfSteps(123)
	}
}

func Benchmark_numberOfSteps_bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberOfSteps_bit(123)
	}
}
