package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func computeOverlap(covers []string) int {
	arr := make([][]int, 1000)
	for i := range arr {
		arr[i] = make([]int, 1000)
	}

	for _, c := range covers {
		bits := strings.Split(c, " ")
		nums := strings.Split(bits[2], ",")
		numX, _ := strconv.Atoi(nums[0])
		numY, _ := strconv.Atoi(nums[1][:len(nums[1])-1])

		lens := strings.Split(bits[3], "x")
		lenX, _ := strconv.Atoi(lens[0])
		lenY, _ := strconv.Atoi(lens[1])

		for i := numX; i < numX+lenX; i++ {
			for j := numY; j < numY+lenY; j++ {
				arr[i][j]++
			}
		}
	}

	counts := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i][j] >= 2 {
				counts++
			}
		}
	}

	return counts
}

func solveDay3Part1(file string) {
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

	fmt.Printf("Day 3 [1]. %v\n", computeOverlap(arr))

}
