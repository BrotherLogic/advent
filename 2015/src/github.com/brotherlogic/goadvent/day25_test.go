package main

import "testing"

func TestDay25P1(t *testing.T) {
	cases := []struct {
		row int
		col int
		val int
	} {
		{1,1,20151125},
		{4,3,21345942},
		{6,6,27995004},
	}

	for _, c := range cases {
		answer := FindCode(20151125, c.row, c.col)
		if answer != c.val {
			t.Errorf("Spec(%v,%v) == %v, want %v", c.row, c.col, answer, c.val)
		}
	}

}
