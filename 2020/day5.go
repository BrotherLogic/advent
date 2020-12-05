package a2020

import (
	"log"
	"sort"
	"strconv"
)

//GetHighestSeat does this
func GetHighestSeat(lines []string) (int, int) {
	best := 0
	founds := []int{}
	for _, line := range lines {
		_, _, seat := findSeat(line)
		founds = append(founds, seat)
		if seat > best {
			best = seat
		}
	}

	mine := 0
	sort.Ints(founds)
	for i, seat := range founds {
		if i < len(founds)-1 && founds[i+1] == seat+2 {
			mine = seat + 1
		}
	}

	return best, mine
}

func locate(loc string, pointer, bot, top int) int {
	if pointer < len(loc) {
		if loc[pointer] == 'F' || loc[pointer] == 'L' {
			return locate(loc, pointer+1, bot, (bot+top)/2)
		} else {
			return locate(loc, pointer+1, (bot+top)/2, top)
		}
	}

	return top
}

func convertBin(loc string) int {
	nStr := ""
	for i := range loc {
		if loc[i] == 'F' || loc[i] == 'L' {
			nStr = nStr + "0"
		} else {
			nStr = nStr + "1"
		}
	}

	val, _ := strconv.ParseInt(nStr, 2, 32)
	return int(val)
}

func findSeat(loc string) (int, int, int) {
	row := locate(string(loc[0:7]), 0, 0, 127)
	row2 := convertBin(string(loc[0:7]))
	if row != row2 {
		log.Fatalf("Bad comp of row: %v and %v -> %v", row, row2, loc[0:7])
	}
	col := locate(string(loc[7:]), 0, 0, 7)
	col2 := convertBin(string(loc[7:]))
	if row != row2 {
		log.Fatalf("Bad comp of col: %v and %v -> %v", col, col2, loc)
	}
	return row, col, row*8 + col
}
