package richest_customer_wealth

import "testing"

func Test_maximumWealth(t *testing.T) {
	tests := []struct {
		name     string
		accounts [][]int
		want     int
	}{
		{
			name:     "all richest",
			accounts: [][]int{{1, 2, 3}, {3, 2, 1}},
			want:     6,
		},
		{
			name:     "one richest",
			accounts: [][]int{{1, 7}, {7, 3}, {3, 5}},
			want:     10,
		},
		{
			name:     "bigger example",
			accounts: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			want:     17,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := maximumWealth(tt.accounts); got != tt.want {
				t.Errorf("maximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
