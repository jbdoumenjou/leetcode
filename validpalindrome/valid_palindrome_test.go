package validpalindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "Example 2",
			s:    "race a car",
			want: false,
		},
		{
			name: "Example 3",
			s:    " ",
			want: true,
		},
		{
			name: "Example 4",
			s:    "aa",
			want: true,
		},
		{
			name: "Example 5",
			s:    "0P0",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_isPalindrome2(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "Example 2",
			s:    "race a car",
			want: false,
		},
		{
			name: "Example 3",
			s:    " ",
			want: true,
		},
		{
			name: "Example 4",
			s:    "aa",
			want: true,
		},
		{
			name: "Example 5",
			s:    "0P0",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_isPalindrome3(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "Example 2",
			s:    "race a car",
			want: false,
		},
		{
			name: "Example 3",
			s:    " ",
			want: true,
		},
		{
			name: "Example 4",
			s:    "aa",
			want: true,
		},
		{
			name: "Example 5",
			s:    "0P0",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome3(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome("0A man, a plan, a canal: Panama0")
	}
}
func Benchmark_isPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome2("0A man, a plan, a canal: Panama0")
	}
}
func Benchmark_isPalindrome3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome3("0A man, a plan, a canal: Panama0")
	}
}
func Benchmark_isPalindrome4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome4("0A man, a plan, a canal: Panama0")
	}
}
func Benchmark_isPalindrome5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome5("0A man, a plan, a canal: Panama0")
	}
}
