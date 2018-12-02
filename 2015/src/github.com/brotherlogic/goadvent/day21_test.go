package main

import "testing"

func TestDay21P1(t *testing.T) {
	one := Fighter{8, 5, 5}
	two := Fighter{12, 7, 2}

	if Fight(one,two) != 1 {
		t.Errorf("Fail on fight\n")
	}
}
