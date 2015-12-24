package main

import "testing"

func TestDay20P1(t *testing.T) {
		cases := []struct {
		in int
		want int
	} {
		{1, 10},
		{2, 30},
		{3, 40},
		{4, 70},
		{5, 60},
		{6, 120},
		{7, 80},
		{8, 150},
		{9, 130},
	}

	for _, c := range cases {
		got := ComputePresents(c.in)
		if got != c.want {
			t.Errorf("Spec(%v) == %v, want %v", c.in, got, c.want)
		}
	}

	
}
