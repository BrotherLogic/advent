package main

import "testing"

var day7Table = []struct {
	in   []string
	in2  int
	out  string
	out2 int
}{
	{[]string{
		"Step C must be finished before step A can begin.",
		"Step C must be finished before step F can begin.",
		"Step A must be finished before step B can begin.",
		"Step A must be finished before step D can begin.",
		"Step B must be finished before step E can begin.",
		"Step D must be finished before step E can begin.",
		"Step F must be finished before step E can begin.",
	}, 2, "CABDFE", 15}}

func TestDay7(t *testing.T) {
	for _, tt := range day7Table {
		order := resolveOrder(tt.in, 6)
		if order != tt.out {
			t.Errorf("Error in order: %v (%v)", order, tt.out)
		}

		timeTaken := resolveTime(tt.in, 6, tt.in2, 0)
		if timeTaken != tt.out2 {
			t.Errorf("Error in time proc: %v (%v)", timeTaken, tt.out2)
		}
	}
}
