package a2020

import "testing"

func TestCount(t *testing.T) {
	lines := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}

	a, b := CountLines(lines)
	if a != 11 {
		t.Errorf("Bad line count: %v", a)
	}

	if b != 6 {
		t.Errorf("Bad line count 2: %v", b)
	}
}
