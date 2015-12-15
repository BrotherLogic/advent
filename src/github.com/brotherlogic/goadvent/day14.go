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

func dayfourteen() {
	matcher := regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)
	file,_ := os.Open("input-day14")
	scanner := bufio.NewScanner(file)
	best := 0
	best_name := ""
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
	}
	fmt.Printf("Best = %v in %v\n", best_name, best)
}
