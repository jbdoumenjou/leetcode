package main

import (
	"reflect"
	"testing"
)

var middleNodeTests = []struct {
	name  string
	input *ListNode
	want  *ListNode
}{
	{
		name: "odd list",
		input: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{Val: 5},
					},
				},
			},
		},
		want: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{Val: 5},
			},
		},
	},
	{
		name: "event list",
		input: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
		want: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: &ListNode{Val: 6},
			},
		},
	},
}

func Test_middleNode(t *testing.T) {
	for _, tt := range middleNodeTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := middleNode(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_middleNode_ptr(t *testing.T) {
	for _, tt := range middleNodeTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := middleNode_ptr(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_middleNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		middleNode(middleNodeTests[0].input)
	}
}

func Benchmark_middleNode_ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		middleNode_ptr(middleNodeTests[0].input)
	}
}
