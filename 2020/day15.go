package a2020

type pair struct {
	first int
	last  int
}

//RunGame runs the game
func RunGame(nums []int, goal int) int {
	m := make(map[int]pair)

	for i, n := range nums {
		m[n] = pair{last: i + 1}
	}

	turn := len(nums) + 1
	last := nums[len(nums)-1]
	for turn <= goal {
		val := m[last]
		//First time we've seen this
		if val.first == 0 {
			last = 0
			val := m[last]
			val.first, val.last = val.last, turn
			m[last] = val
		} else {
			last = val.last - val.first
			val = m[last]
			val.first, val.last = val.last, turn
			m[last] = val
		}

		turn++
	}
	return last
}
