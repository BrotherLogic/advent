package main

import "testing"

func TestDay2(t *testing.T) {

	cases := []struct {
		in string
		want int
	} {
		{"2x3x4", 58},
		{"1x1x10", 43},
	}

	for _, c := range cases {
		got := ComputeAmountOfPaper(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
