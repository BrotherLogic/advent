package a2019

import "testing"

func TestRunOrbit(t *testing.T) {
	mapper := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}

	result, span := RunOrbits(mapper)
	if result != 54 {
		t.Errorf("Bad orbits: %v", result)
	}

	if span != 4 {
		t.Errorf("Bad span: %v", span)
	}
}
