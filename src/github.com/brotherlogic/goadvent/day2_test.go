package main

import "testing"

func TestDay2P2(t *testing.T) {
	cases := []struct {
		in string
		want int
	} {
		{"2x3x4", 34},
		{"1x1x10", 14},
	}

	for _, c := range cases {
		got := ComputeAmountOfRibbon(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

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
