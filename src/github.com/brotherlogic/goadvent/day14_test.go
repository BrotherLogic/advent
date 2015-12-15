package main

import "testing"

func TestDay14P1(t *testing.T) {
	cases := []struct {
		speed int
		stime int
		wait int
		time int
		distance int
	} {
		{14,10,127,1,14},
		{14,10,127,10,140},
		{14,10,127,11,140},
		{14,10,127,1000,1120},
		{16,11,162,1,16},
		{16,11,162,10,160},
		{16,11,162,11,176},
		{16,11,162,1000,1056},
	}

	for _, c := range cases {
		answer := ComputeDistance(c.speed,c.stime,c.wait,c.time)
		if answer != c.distance {
			t.Errorf("Spec(%v,%v) == %v, want %v", c.speed, c.time, answer, c.distance)
		}
	}
}
