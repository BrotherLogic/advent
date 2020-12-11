package a2020

//FindFirst number that fails
func FindFirst(nums []int64, preceed int) int64 {
	for check := preceed; check < len(nums); check++ {
		passed := false
		for i := check - 1; i > check-preceed-1; i-- {
			for j := i - 1; j > check-preceed-1; j-- {
				if nums[i]+nums[j] == nums[check] {
					passed = true
				}
			}
		}

		if !passed {
			return nums[check]
		}
	}

	return 0
}

func sum(nums []int64) int64 {
	count := int64(0)
	for _, num := range nums {
		count += num
	}

	return count
}

func max(nums []int64) int64 {
	best := nums[0]
	for _, num := range nums {
		if num > best {
			best = num
		}
	}
	return best
}

func min(nums []int64) int64 {
	best := nums[0]
	for _, num := range nums {
		if num < best {
			best = num
		}
	}
	return best
}

//FindContiguous sum of numbers
func FindContiguous(nums []int64, val int64) (int64, int64) {
	for head := 0; head < len(nums); head++ {
		for tail := head + 1; tail < len(nums); tail++ {
			sumv := sum(nums[head:tail])
			if sumv == val {
				return min(nums[head:tail]), max(nums[head:tail])
			}
		}
	}

	return 0, 0
}
