package a2020

import (
	"log"
	"sort"
)

func maxi(nums []int) int {
	best := nums[0]
	for _, num := range nums {
		if num > best {
			best = num
		}
	}
	return best
}

func minint(a, b int) int {
	if a < b {
		return a
	}

	return b
}

//Travel part 2
func Travel(nums []int) int {
	cp := make([]int, len(nums))
	copy(cp, nums)

	cp = append(cp, 0)
	cp = append(cp, maxi(cp)+3)

	sort.Ints(cp)

	magic := make([]int, len(cp))
	magic[len(magic)-1] = 1
	for i := len(magic) - 2; i >= 0; i-- {
		sum := 0
		for j := i + 1; j < minint(i+4, len(cp)); j++ {
			if cp[j] <= cp[i]+3 {
				sum += magic[j]
			}
		}
		magic[i] = sum
	}

	return magic[0]
}

func compute(nums []int, pointer int) int {
	log.Printf("COMPUTE %v / %v", pointer, len(nums))
	count := 0
	for i := pointer + 1; i < len(nums); i++ {
		if nums[i]-nums[pointer] <= 3 {
			count += compute(nums, i)
		}
	}

	if pointer == len(nums)-1 {
		return 1
	}
	return count
}

//FindDiff joltage
func FindDiff(nums []int, val1, val2 int) int {
	cp := make([]int, len(nums))
	copy(cp, nums)

	cp = append(cp, 0)
	cp = append(cp, maxi(cp)+3)

	sort.Ints(cp)

	diffs := make(map[int]int)
	for i := 1; i < len(cp); i++ {
		diffs[cp[i]-cp[i-1]]++
	}

	return diffs[val1] * diffs[val2]
}
