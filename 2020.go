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

func run2020day5() {
	lines := readLines("data/2020-5-1.txt")
	v1, mine := a2020.GetHighestSeat(lines)
	fmt.Printf("Day5-1: %v\n", v1)
	fmt.Printf("Day5-2: %v\n", mine)
}

func run2020day6() {
	lines := readLines("data/2020-6-1.txt")
	v1, v2 := a2020.CountLines(lines)
	fmt.Printf("Day6-1: %v\n", v1)
	fmt.Printf("Day6-2: %v\n", v2)
}

func run2020day7() {
	lines := readLines("data/2020-7-1.txt")
	count := a2020.Derive(lines, "shiny gold")
	fmt.Printf("Day7-1: %v\n", count)
	deriv := a2020.Delve(lines, "shiny gold")
	fmt.Printf("Day7-2: %v\n", deriv)
}

func run2020day8() {
	code := readLines("data/2020-8-1.txt")
	acc := a2020.RunCode(code)
	fmt.Printf("Day8-1: %v\n", acc)
	acc = a2020.FixAndRunCode(code)
	fmt.Printf("Day8-2: %v\n", acc)
}

func run2020day9() {
	lines := loadBigNums("data/2020-9-1.txt")
	first := a2020.FindFirst(lines, 25)
	fmt.Printf("Day9-1: %v\n", first)
	contMin, contMax := a2020.FindContiguous(lines, first)
	fmt.Printf("Day9-2: %v\n", contMin+contMax)
}

func run2020day10() {
	lines := loadNums("data/2020-10-1.txt")
	found := a2020.FindDiff(lines, 1, 3)
	fmt.Printf("Day10-1: %v\n", found)
	travel := a2020.Travel(lines)
	fmt.Printf("Day10-2: %v\n", travel)
}

func run2020day11() {
	lines := readLines("data/2020-11-1.txt")
	seats := a2020.CountAndRunSeats(lines, 0)
	fmt.Printf("Day11-1: %v\n", seats)
	seats2 := a2020.CountAndRunSeats(lines, 1)
	fmt.Printf("Day11-2: %v\n", seats2)
}

func run2020day12() {
	lines := readLines("data/2020-12-1.txt")
	dist := a2020.MakeMoves(lines)
	fmt.Printf("Day12-1: %v\n", dist)
	wp := a2020.MoveWaypoint(lines)
	fmt.Printf("Day12-2: %v\n", wp)
}

func run2020day13() {
	lines := readLines("data/2020-13-1.txt")
	tim := a2020.GetEarliestTime(lines)
	fmt.Printf("Day13-1: %v\n", tim)
	t2 := a2020.GetAlignment(lines[1])
	fmt.Printf("Day13-2: %v\n", t2)
}

func run2020day14() {
	lines := readLines("data/2020-14-1.txt")
	tim := a2020.RunMask(lines)
	fmt.Printf("Day14-1: %v\n", tim)
	tim2 := a2020.RunMask2(lines)
	fmt.Printf("Day14-2: %v\n", tim2)
}

func run2020day15() {
	res := a2020.RunGame([]int{0, 14, 6, 20, 1, 4}, 2020)
	fmt.Printf("Day15-1: %v\n", res)
	res2 := a2020.RunGame([]int{0, 14, 6, 20, 1, 4}, 30000000)
	fmt.Printf("Day15-2: %v\n", res2)
}

func run2020day16() {
	lines := readLines("data/2020-16-1.txt")
	res, _ := a2020.FindTicket(lines)
	fmt.Printf("Day16-1: %v\n", res)
	res2 := a2020.BuildTicket(lines, "depart")
	fmt.Printf("Day16-2: %v\n", res2)
}

func run2020day17() {
	lines := readLines("data/2020-17-1.txt")
	res := a2020.RunLife(lines, 6)
	fmt.Printf("Day17-1: %v\n", res)
	res2 := a2020.RunLife2(lines, 6)
	fmt.Printf("Day17-2: %v\n", res2)
}

func run2020day18() {
	lines := readLines("data/2020-18-1.txt")
	res := a2020.RunSum(lines, map[string]int{"*": 1, "+": 1, "(": 0, ")": 0})
	fmt.Printf("Day18-1: %v\n", res)
	res2 := a2020.RunSum(lines, map[string]int{"*": 1, "+": 2, "(": 0, ")": 0})
	fmt.Printf("Day18-2: %v\n", res2)
}
