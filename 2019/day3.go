package a2019

import (
	"strconv"
)

func fill(mapper map[int]map[int]int, path []string) {
	pointx := 1
	pointy := 1

	counts := 0
	for _, elem := range path {
		adjusterx := 0
		adjustery := 0

		if elem[0] == 'U' {
			adjustery = 1
		} else if elem[0] == 'D' {
			adjustery = -1
		} else if elem[0] == 'R' {
			adjusterx = 1
		} else if elem[0] == 'L' {
			adjusterx = -1
		}

		num, _ := strconv.ParseInt(elem[1:], 10, 32)
		for i := 0; i < int(num); i++ {
			counts++

			pointx += adjusterx
			pointy += adjustery

			base, ok := mapper[pointx]
			if !ok {
				mapper[pointx] = make(map[int]int)
				base = mapper[pointx]
			}
			base[pointy] = counts
		}
	}
}

func find(mapper map[int]map[int]int, path []string) ([]int, []int) {
	pointx := 1
	pointy := 1

	pointsx := []int{}
	pointsy := []int{}

	for _, elem := range path {
		adjusterx := 0
		adjustery := 0

		if elem[0] == 'U' {
			adjustery = 1
		} else if elem[0] == 'D' {
			adjustery = -1
		} else if elem[0] == 'R' {
			adjusterx = 1
		} else if elem[0] == 'L' {
			adjusterx = -1
		}

		num, _ := strconv.ParseInt(elem[1:], 10, 32)
		for i := 0; i < int(num); i++ {
			pointx += adjusterx
			pointy += adjustery

			if mapper[pointx][pointy] > 0 {
				pointsx = append(pointsx, pointx)
				pointsy = append(pointsy, pointy)
			}
		}
	}

	return pointsx, pointsy
}

func findLen(mapper map[int]map[int]int, path []string) []int {
	pointx := 1
	pointy := 1

	count := 0
	lens := []int{}
	for _, elem := range path {
		adjusterx := 0
		adjustery := 0

		if elem[0] == 'U' {
			adjustery = 1
		} else if elem[0] == 'D' {
			adjustery = -1
		} else if elem[0] == 'R' {
			adjusterx = 1
		} else if elem[0] == 'L' {
			adjusterx = -1
		}

		num, _ := strconv.ParseInt(elem[1:], 10, 32)
		for i := 0; i < int(num); i++ {
			count++
			pointx += adjusterx
			pointy += adjustery

			if mapper[pointx][pointy] > 0 {
				lens = append(lens, mapper[pointx][pointy]+count)
			}
		}
	}

	return lens
}

func abs(val int) int {
	if val < 0 {
		return 0 - val
	}
	return val
}

//ComputeDistance runs day3 calc
func ComputeDistance(path1, path2 []string) int {
	drawArr := make(map[int]map[int]int)

	fill(drawArr, path1)
	pointsx, pointsy := find(drawArr, path2)

	maxPath := 99999
	for i := range pointsx {
		val := abs(pointsx[i]-1) + abs(pointsy[i]-1)
		if val < maxPath {
			maxPath = val
		}
	}

	return maxPath
}

//ComputeLen of path
func ComputeLen(path1, path2 []string) int {
	drawArr := make(map[int]map[int]int)

	fill(drawArr, path1)
	lens := findLen(drawArr, path2)

	maxPath := 99999
	for _, val := range lens {
		if val < maxPath {
			maxPath = val
		}
	}

	return maxPath
}
