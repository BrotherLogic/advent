package a2020

import (
	"testing"
)

func TestTime(t *testing.T) {
	lines := []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}

	tim := GetEarliestTime(lines)
	if tim != 295 {
		t.Errorf("Bad time: %v", tim)
	}

	t2 := GetAlignment("1789,37,47,1889")
	if t2 != 1202161486 {
		t.Errorf("Bad align: %v", t2)
	}
}
