package main

import "testing"

var day5Part1Table = []struct {
	in  string
	out int
}{
	{"dabAcCaCBAcCcaDA", 10},
}

func TestDay5Part1(t *testing.T) {
	for _, tt := range day5Part1Table {
		if len(collapse(tt.in)) != tt.out {
			t.Errorf("Error in collapse: %v", collapse(tt.in))
		}
	}

}
