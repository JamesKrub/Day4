package removeduplicate_test

import (
	"KBTGCourse/day4/removeduplicate"
	"testing"
)

func TestDuplicateString(t *testing.T) {
	var flagtests = []struct {
		input  []int
		output []int
	}{ // input                         output
		{[]int{1, 1, 2, 3, 4, 6, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{0, 1, 2, 3, 3, 7, 8, 9}, []int{0, 1, 2, 3, 7, 8, 9}},
		{[]int{5, 6, 8, 7, 9, 10, 4}, []int{4, 5, 6, 7, 8, 9, 10}},
	}
	for _, tt := range flagtests {
		r := removeduplicate.RemoveDuplicate(tt.input)
		if len(r) != len(tt.output) {
			t.Errorf("Expect %v got %v", len(r), len(tt.output))
		}
		for i := 0; i < len(r); i++ {
			if r[i] != tt.output[i] {
				t.Errorf("Expect %v got %v", r[i], tt.output[i])
			}
		}
	}
}

func BenchmarkDuplicateString(b *testing.B) {
	var flagtests = []struct {
		input  []int
		output []int
	}{ // input                         output
		{[]int{1, 1, 2, 3, 4, 6, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{0, 1, 2, 3, 3, 7, 8, 9}, []int{0, 1, 2, 3, 7, 8, 9}},
		{[]int{5, 6, 8, 7, 9, 10, 4}, []int{4, 5, 6, 7, 8, 9, 10}},
	}
	for _, tt := range flagtests {
		r := removeduplicate.RemoveDuplicate(tt.input)
		if len(r) != len(tt.output) {
			b.Errorf("Expect %v got %v", len(r), len(tt.output))
		}
		for i := 0; i < len(r); i++ {
			if r[i] != tt.output[i] {
				b.Errorf("Expect %v got %v", r[i], tt.output[i])
			}
		}
	}
}
