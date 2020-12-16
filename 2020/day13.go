package a2020

import (
	"strconv"
	"strings"
)

func getNumbers(line string) []int64 {
	numstr := strings.Split(line, ",")
	nums := []int64{}

	for _, sn := range numstr {
		if sn != "x" {
			nv, _ := strconv.ParseInt(sn, 10, 64)
			nums = append(nums, nv)
		}
	}

	return nums
}

func prod(nums []int64) int64 {
	val := nums[0]
	for i := 1; i < len(nums); i++ {
		val *= nums[i]
	}
	return val
}

func findValue(base, multiplier, number, adj int64) (int64, int64) {
	x := base
	multi := int64(-1)
	for true {
		if x%number == adj%number && multi > 0 {
			return multi, x - multi
		}

		if x%number == adj%number {
			multi = x
		}
		x += multiplier
	}

	return -1, -1
}

//GetAlignment with things
func GetAlignment(line string) int64 {
	base := int64(0)
	multiplier := int64(1)
	count := len(strings.Split(line, ","))

	for i, num := range strings.Split(line, ",") {
		if num != "x" {
			num, _ := strconv.ParseInt(num, 10, 64)
			nbase, nmulti := findValue(base, multiplier, num, int64(count-i))
			base, multiplier = nbase, nmulti
		}
	}

	return base - int64(count)
}

//GetEarliestTime sol to puzzle 13
func GetEarliestTime(lines []string) int64 {

	ts, _ := strconv.ParseInt(lines[0], 10, 64)
	nums := getNumbers(lines[1])

	bestTime := int64(99999999999)
	bestID := int64(0)
	for _, num := range nums {
		val := ts / num
		ntime := (val + 1) * num
		if ntime < bestTime {
			bestTime = ntime
			bestID = num
		}
	}

	return (bestTime - ts) * bestID
}
