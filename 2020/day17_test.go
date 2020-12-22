package a2020

import "testing"

func TestLife1(t *testing.T) {
	lines := []string{
		".#.",
		"..#",
		"###",
	}

	res := RunLife(lines, 6)
	if res != 112 {
		t.Errorf("Bad run of lie: %v", res)
	}

	res2 := RunLife2(lines, 6)
	if res2 != 848 {
		t.Errorf("Bad run: %v", res2)
	}
}
