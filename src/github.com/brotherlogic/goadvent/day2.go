package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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


func daytwo() {
	total :=0
	
	file, err := os.Open("input-day2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += ComputeAmountOfPaper(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total = %d\n", total)
}
