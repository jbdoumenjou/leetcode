package ransomnote

import "testing"

var canConstructTests = []struct {
	name       string
	ransomNote string
	magazine   string
	want       bool
}{
	{
		name:       "example 1",
		ransomNote: "a",
		magazine:   "b",
		want:       false,
	},
	{
		name:       "example 2",
		ransomNote: "aa",
		magazine:   "ab",
		want:       false,
	},
	{
		name:       "example 3",
		ransomNote: "aa",
		magazine:   "aab",
		want:       true,
	},
	{
		name:       "example 4",
		ransomNote: "bg",
		magazine:   "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj",
		want:       true,
	},
	{
		name:       "example 5",
		ransomNote: "aab",
		magazine:   "baa",
		want:       true,
	},
	{
		name:       "example 6",
		ransomNote: "fffbfg",
		magazine:   "effjfggbffjdgbjjhhdegh",
		want:       true,
	},
}

func Test_canConstruct(t *testing.T) {
	for _, tt := range canConstructTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.ransomNote, tt.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canConstruct_array(t *testing.T) {
	for _, tt := range canConstructTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct_array(tt.ransomNote, tt.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_canConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canConstruct(canConstructTests[0].ransomNote, canConstructTests[0].magazine)
	}
}

func Benchmark_canConstruct_array(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canConstruct_array(canConstructTests[0].ransomNote, canConstructTests[0].magazine)
	}
}
