package main

import "testing"

var day6Table = []struct {
	in   []string
	maxV int
	out  int
	out2 int
}{
	{[]string{"1, 1",
		"1, 6",
		"8, 3",
		"3, 4",
		"5, 5",
		"8, 9",
	}, 32, 17, 16}}

func TestDay6(t *testing.T) {
	for _, tt := range day6Table {
		maxSize, maxRegion := computeMaxSize(tt.in, tt.maxV)
		if maxSize != tt.out {
			t.Errorf("Error in compute: %v (%v)", maxSize, tt.out)
		}
		if maxRegion != tt.out2 {
			t.Errorf("Error in second compute: %v", maxRegion)
		}
	}
}
