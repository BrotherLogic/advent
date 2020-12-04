package a2020

import "testing"

func TestExample(t *testing.T) {
	lines := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	trees := CountTrees(lines, 0, 0, 3, 1)
	if trees != 7 {
		t.Errorf("Bad path: %v", trees)
	}

	treesAll := CountAllSlopes(lines)
	if treesAll != 336 {
		t.Errorf("Bad pathss: %v", treesAll)
	}
}
