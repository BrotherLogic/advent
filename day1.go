package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(arr []int) int {
	val := 0
	for _, v := range arr {
		val += v
	}
	return val
}

func lookForTwo(arr []int) int {
	counts := make(map[int]int)
	counts[0] = 1
	cVal := 0
	for true {
		for _, v := range arr {
			cVal += v
			if counts[cVal] == 1 {
				return cVal
			} else {
				counts[cVal]++
			}
		}
	}

	return -1
}

func solveDay1Part1(file string) {
	arr := []int{}
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 1 [1]. %v\n", sum(arr))
	fmt.Printf("Day 1 [2]. %v\n", lookForTwo(arr))
}
