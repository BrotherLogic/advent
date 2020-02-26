package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type guard struct {
	timestamp int64
	action    string
}

func convertTime(t string) int64 {
	tim, err := time.Parse("2006-01-02 15:04", strings.Replace(t, "1518", "2018", -1))
	if err != nil {
		log.Fatalf("Parse error: %v", err)
	}

	return tim.Unix()
}

func computeSleep(arr []string) (int, int) {
	gArr := []guard{}

	for _, a := range arr {
		elems := strings.Split(a, "]")
		g := guard{timestamp: convertTime(elems[0][1:]), action: elems[1]}
		gArr = append(gArr, g)
	}

	sort.SliceStable(gArr, func(i, j int) bool { return gArr[i].timestamp < gArr[j].timestamp })

	gMap := make(map[string]int64)
	mMap := make(map[string][]int)
	guardName := ""
	currStart := int64(0)
	for _, g := range gArr {
		if strings.Contains(g.action, "#") {
			guardName = g.action
			if _, ok := mMap[guardName]; !ok {
				mMap[guardName] = make([]int, 60)
			}
		} else if strings.Contains(g.action, "asleep") {
			currStart = g.timestamp
		} else {
			gMap[guardName] += (g.timestamp - currStart) / 60
			for min := currStart; min < g.timestamp; min += 60 {
				minute := time.Unix(min, 0).Minute()
				mMap[guardName][minute]++
			}
		}
	}

	maxGV := int64(0)
	maxGN := ""
	for name, val := range gMap {
		if val > maxGV {
			maxGV = val
			maxGN = name
		}
	}

	maxM := 0
	maxV := 0
	for i, val := range mMap[maxGN] {
		if val > maxV {
			maxV = val
			maxM = i
		}
	}

	maxGA := ""
	maxNum := 0
	maxMin := 0
	for g, val := range mMap {
		for i, v := range val {
			if v > maxNum {
				maxGA = g
				maxMin = i
				maxNum = v
			}
		}
	}

	gn := strings.Fields(maxGN)
	gnum, _ := strconv.Atoi(gn[1][1:])

	gn2 := strings.Fields(maxGA)
	gnum2, _ := strconv.Atoi(gn2[1][1:])

	return gnum * maxM, gnum2 * maxMin
}

func solveDay4Part1(file string) {
	arr := []string{}
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	v1, v2 := computeSleep(arr)
	fmt.Printf("Day 4 [1]. %v\n", v1)
	fmt.Printf("Day 4 [2]. %v\n", v2)
}
