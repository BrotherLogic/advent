package main

import "testing"

var day5Part1Table = []struct {
	in   string
	out  int
	out2 int
}{
	{"dabAcCaCBAcCcaDA", 10, 4},
}

func TestDay5Part1(t *testing.T) {
	for _, tt := range day5Part1Table {
		col := len(collapse(tt.in))
		short := shortest(tt.in)
		if col != tt.out || short != tt.out2 {
			t.Errorf("Error in collapse: %v and %v", col, short)
		}
	}

}
