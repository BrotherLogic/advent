package main

import "testing"

func TestDay3P1(t *testing.T) {

	cases := []struct {
		in string
		want int
	} {
		{">", 0},
		{"^>v<", 0},
		{"^v^v^v^v^v", 0},
	}

	for _, c := range cases {
		got := ComputeNumberOfHouses(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
