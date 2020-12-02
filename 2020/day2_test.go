package a2020

import "testing"

func TestParser(t *testing.T) {
	lines := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	vc1, vc2 := CountValid(lines)
	if vc1 != 2 {
		t.Errorf("Bad Valid: %v", vc1)
	}

	if vc2 != 1 {
		t.Errorf("Bad new valid: %v", vc2)
	}
}
