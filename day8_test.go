package main

import "testing"

var day8Table = []struct {
	in   []int
	out  int
	out2 int
}{
	{[]int{
		2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2,
	}, 138, 66}}

func TestDay8(t *testing.T) {
	for _, tt := range day8Table {
		sumv, sumr := computeSum(tt.in)
		if sumv != tt.out {
			t.Errorf("Pah %v (%v)", sumv, tt.out)
		}
		if sumr != tt.out2 {
			t.Errorf("Pah %v (%v)", sumr, tt.out2)
		}
	}
}
