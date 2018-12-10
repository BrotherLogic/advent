package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	GRID = 500
)

type gridpiece struct {
	identifier int
	distance   int
	sumdist    int
}

func abs(v int) int {
	if v < 0 {
		return 0 - v
	}

	return v
}

func computeDist(i, j, x, y int) int {
	return abs(i-x) + abs(j-y)
}

func computeMaxSize(coords []string, maxD int) (int, int) {
	// Initialise the grid
	city := make([][]gridpiece, GRID)
	for i := range city {
		city[i] = make([]gridpiece, GRID)
		for j := range city {
			city[i][j] = gridpiece{-1, GRID * 20, 0}
		}
	}

	for ci, coord := range coords {
		elems := strings.Split(coord, ",")
		x, _ := strconv.Atoi(elems[0])
		y, _ := strconv.Atoi(elems[1][1:])

		for i := range city {
			for j := range city[i] {
				dist := computeDist(i, j, x, y)
				city[i][j].sumdist += dist
				if dist < city[i][j].distance {
					city[i][j].distance = dist
					city[i][j].identifier = ci
				} else if dist == city[i][j].distance {
					city[i][j].identifier = -1
				}
			}
		}
	}

	blacklisted := make(map[int]bool)
	sizes := make(map[int]int)
	count := 0
	for i := range city {
		for j := range city {
			if i == 0 || i == GRID-1 || j == 0 || j == GRID-1 {
				blacklisted[city[i][j].identifier] = true
			} else {
				sizes[city[i][j].identifier]++
			}
			if city[i][j].sumdist < maxD {
				count++
			}
		}
	}

	maxSize := 0
	for ident, size := range sizes {
		if !blacklisted[ident] && size > maxSize {
			maxSize = size
		}
	}

	return maxSize, count
}

func solveDay6(file string) {
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

	v1, v2 := computeMaxSize(arr, 10000)
	fmt.Printf("Day 6 [1]. %v\n", v1)
	fmt.Printf("Day 6 [2]. %v\n", v2)
}
