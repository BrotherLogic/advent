package main

import "testing"

func TestDay4P1(t *testing.T) {

	cases := []struct {
		in string
		want int
	} {
		{"abcdef", 0},
		{"pqrstuv", 0},
	}

	for _, c := range cases {
		got := SolveHash(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}