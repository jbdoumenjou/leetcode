package main

import (
	"reflect"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "One Fizz",
			n:    3,
			want: []string{"1", "2", "Fizz"},
		},
		{
			name: "Fizz Buzz",
			n:    5,
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "Fizz Buzz FizzBuzz",
			n:    15,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := fizzBuzz(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
