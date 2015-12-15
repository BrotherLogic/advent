package main

import "bufio"
import "fmt"
import "os"
import "regexp"
import "strconv"

func ComputeDistance(speed int, stime int, wait int, time int) int {
	maxdist := speed*stime
	period_time := stime+wait
	
	period_dist := (time/period_time)*maxdist
	period_extra := time%period_time

	if period_extra > stime {
		return period_dist + speed * stime
	} else {
		return period_dist + speed*period_extra
	}
}

func ComputeDistances(speed int, stime int, wait int, time int) []int {
	var distances = make([]int, time)
	for i := 0 ; i < len(distances) ; i++ {
		distances[i] = ComputeDistance(speed, stime, wait, i+1)
	}
	return distances
}

func dayfourteen() {
	matcher := regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)
	file,_ := os.Open("input-day14")
	scanner := bufio.NewScanner(file)
	best := 0
	best_name := ""
	var m map[string][]int
	m = make(map[string][]int)
	for scanner.Scan() {
		text := scanner.Text()
		match := matcher.FindAllStringSubmatch(text,-1)
		speed,_ := strconv.Atoi(match[0][2])
		stime,_ := strconv.Atoi(match[0][3])
		rest,_ := strconv.Atoi(match[0][4])
		dist := ComputeDistance(speed,stime,rest, 2503)
		if dist > best {
			best = dist
			best_name = match[0][1]
		}
		m[match[0][1]] = ComputeDistances(speed,stime,rest,2503)
	}
	fmt.Printf("Best = %v in %v\n", best_name, best)

	var scores map[string]int
	scores = make(map[string]int)
	for i := 0; i < 2503; i++ {
		best := 0
		best_name := ""
		for key, value := range m {
			if value[i] > best {
				best = value[i]
				best_name = key
			}
		}
		scores[best_name] += 1
	}

	best = 0
	best_score := ""
	for key, value := range scores {
		if value > best {
			best = value
			best_score = key
		}
	}
	fmt.Printf("Best Score = %v in %v\n", best, best_score)
}
