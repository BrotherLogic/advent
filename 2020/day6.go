package a2020

//CountLines counts yes answers
func CountLines(lines []string) (int, int) {
	counts := []int{}
	countsall := []int{}
	seen := make(map[string]int)
	gc := 0
	for _, line := range lines {
		if line == "" {
			counts = append(counts, len(seen))
			count2 := 0
			for _, val := range seen {
				if val == gc {
					count2++
				}
			}
			countsall = append(countsall, count2)
			seen = make(map[string]int)
			gc = 0
		} else {
			for _, el := range line {
				seen[string(el)]++
			}
			gc++
		}
	}
	counts = append(counts, len(seen))
	count2 := 0
	for _, val := range seen {
		if val == gc {
			count2++
		}
	}
	countsall = append(countsall, count2)

	count := 0
	count2 = 0
	for i, c := range counts {
		count += c
		count2 += countsall[i]
	}
	return count, count2
}
