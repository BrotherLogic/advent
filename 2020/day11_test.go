package a2020

import "testing"

func TestSeats(t *testing.T) {
	lines := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}

	res := CountAndRunSeats(lines, 0)
	if res != 37 {
		t.Errorf("Bad seats: %v", res)
	}

	res2 := CountAndRunSeats(lines, 1)
	if res2 != 26 {
		t.Errorf("Bad seats2: %v", res2)
	}
}
