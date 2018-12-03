package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func computeDiff(arr []string) string {
	bestLeft := ""
	bestRight := ""
	bestDiff := 100

	for i, left := range arr {
		for _, right := range arr[i+1:] {
			diff := 0
			for i := 0; i < len(left); i++ {
				if left[i] != right[i] {
					diff++
				}
			}

			if diff < bestDiff {
				bestDiff = diff
				bestLeft = left
				bestRight = right
			}
		}
	}

	superString := ""
	for i := range bestLeft {
		if bestLeft[i] == bestRight[i] {
			superString += string(bestLeft[i])
		}
	}

	return superString
}

func computeChecksum(arr []string) int {
	numTwos := 0
	numThrees := 0

	for _, a := range arr {
		counts := make(map[rune]int)
		for _, c := range a {
			counts[c]++
		}

		done2 := false
		done3 := false
		for _, v := range counts {
			if v == 2 && !done2 {
				numTwos++
				done2 = true
			} else if v == 3 && !done3 {
				numThrees++
				done3 = true
			}
		}
	}

	return numTwos * numThrees
}

func solveDay2(file string) {
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

	fmt.Printf("Day2 [1]. %v\n", computeChecksum(arr))
	fmt.Printf("Day2 [2]. %v\n", computeDiff(arr))
}
