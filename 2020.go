package main

import (
	"fmt"

	a2020 "github.com/brotherlogic/advent/2020"
)

func run2020day1() {
	numbers := loadNums("data/2020-1-1.txt")
	fmt.Printf("Day1-1: %v\n", a2020.ComputeExpenseReport(numbers))
	fmt.Printf("Day1-2: %v\n", a2020.ComputeExpenseReportThree(numbers))
}

func run2020day2() {
	lines := readLines("data/2020-2-1.txt")
	v1, v2 := a2020.CountValid(lines)
	fmt.Printf("Day2-1: %v\n", v1)
	fmt.Printf("Day2-2: %v\n", v2)
}

func run2020day3() {
	lines := readLines("data/2020-3-1.txt")
	v1 := a2020.CountTrees(lines, 0, 0, 3, 1)
	fmt.Printf("Day3-1: %v\n", v1)
	v2 := a2020.CountAllSlopes(lines)
	fmt.Printf("Day3-2: %v\n", v2)
}

func run2020day4() {
	lines := readLines("data/2020-4-1.txt")
	v1 := a2020.CountValid1(lines)
	fmt.Printf("Day4-1: %v\n", v1)
	v2 := a2020.CountValid2(lines)
	fmt.Printf("Day4-2: %v\n", v2)
}
