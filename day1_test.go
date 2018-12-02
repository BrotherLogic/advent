package main

import (
	"testing"
)

var testTable = []struct {
	in  []int
	out int
}{
	{[]int{+1, -2, +3, +1}, 3},
	{[]int{+1, +1, +1}, 3},
	{[]int{+1, +1, -2}, 0},
	{[]int{-1, -2, -3}, -6},
}

var testTablePart2 = []struct {
	in  []int
	out int
}{
	{[]int{+1, -2, +3, +1}, 2},
	{[]int{+1, -1}, 0},
	{[]int{+3, +3, +4, -2, -4}, 10},
	{[]int{-6, +3, +8, +5, -6}, 5},
	{[]int{+7, +7, -2, -7, -4}, 14},
}

func TestRun(t *testing.T) {
	for _, tt := range testTable {
		if sum(tt.in) != tt.out {
			t.Errorf("%v lead to %v, should have been %v", tt.in, sum(tt.in), tt.out)
		}
	}
}

func TestRunPart2(t *testing.T) {
	for _, tt := range testTablePart2 {
		if lookForTwo(tt.in) != tt.out {
			t.Errorf("%v lead to %v in part 2, should have been %v", tt.in, lookForTwo(tt.in), tt.out)
		}
	}
}
