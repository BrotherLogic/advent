package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"strconv"
)

func Min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func ComputeAmountOfPaper(str string) int {
	strs := strings.Split(str,"x")
	length, err := strconv.Atoi(strs[0])
	if err == nil {		
		width, err := strconv.Atoi(strs[1])
		if err == nil {
			height, err := strconv.Atoi(strs[2])
			if err == nil {
				surface1 := 2 * length * width
				surface2 := 2 * width * height
				surface3 := 2 * height * length
				slack := Min(surface1, Min(surface2, surface3)) / 2
				return surface1 + surface2 + surface3 + slack
			}
		}
	} 

	return 0
}

func ComputeAmountOfRibbon(str string) int {
	strs := strings.Split(str,"x")
	length, err := strconv.Atoi(strs[0])
	if err == nil {		
		width, err := strconv.Atoi(strs[1])
		if err == nil {
			height, err := strconv.Atoi(strs[2])
			if err == nil {
				arr := []int {length, width, height}
				sort.Ints(arr)
				return arr[0]*2+arr[1]*2+length*width*height
			}
		}
	} 

	return 0
}


func daytwo() {
	total := 0
	ribbon := 0
	
	file, err := os.Open("input-day2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		total += ComputeAmountOfPaper(text)
		ribbon += ComputeAmountOfRibbon(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Paper = %d\n", total)
	fmt.Printf("Total Ribbon = %d\n", ribbon)
}
