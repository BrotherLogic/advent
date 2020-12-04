package a2020

//CountTrees does what is says on the tin
func CountTrees(lines []string, x, y, right, down int) int {
	if y < len(lines) {
		if lines[y][x] == '#' {
			return 1 + CountTrees(lines, (x+right)%len(lines[0]), (y+down), right, down)
		}
		return CountTrees(lines, (x+right)%len(lines[0]), (y + down), right, down)
	}

	return 0
}

// CountAllSlopes counts all the slopes
func CountAllSlopes(lines []string) int {
	value := 1
	for i, right := range []int{1, 3, 5, 7, 1} {
		down := 1
		if i > 0 && right == 1 {
			down = 2
		}
		value *= CountTrees(lines, 0, 0, right, down)
	}

	return value
}
