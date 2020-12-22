package a2020

import "testing"

func TestTicket(t *testing.T) {
	lines := []string{
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",
		"",
		"your ticket:",
		"7,1,14",
		"",
		"nearby tickets:",
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	}

	scan, _ := FindTicket(lines)
	if scan != 71 {
		t.Errorf("Bad ticket: %v", scan)
	}
}

func TestBuild(t *testing.T) {
	lines := []string{
		"class: 0-1 or 4-19",
		"row: 0-5 or 8-19",
		"seat: 0-13 or 16-19",
		"",
		"your ticket:",
		"11,12,13",
		"",
		"nearby tickets:",
		"3,9,18",
		"15,1,5",
		"5,14,9",
	}
	val := BuildTicket(lines, "")
	if val != 12*11*13 {
		t.Errorf("Bad: %v", val)
	}
}
