package a2020

import (
	"strconv"
	"strings"
)

func isValid(line string) (bool, bool) {
	elems := strings.Fields(line)
	parts := strings.Split(elems[0], "-")

	min, _ := strconv.ParseInt(parts[0], 10, 32)
	max, _ := strconv.ParseInt(parts[1], 10, 32)

	letter := rune(elems[1][0])

	count := int64(0)
	sum := 0
	for _, char := range elems[2] {
		if char == letter {
			count++
		}
	}

	if rune(elems[2][int(min)-1]) == letter {
		sum++
	}
	if rune(elems[2][int(max)-1]) == letter {
		sum++
	}

	return count >= min && count <= max, sum == 1
}

//CountValid the number of vald passwords
func CountValid(lines []string) (int, int) {
	count := 0
	count2 := 0
	for _, line := range lines {
		v1, v2 := isValid(line)
		if v1 {
			count++
		}
		if v2 {
			count2++
		}
	}
	return count, count2
}
