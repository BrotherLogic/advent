package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	a2019 "github.com/brotherlogic/advent/2019"
)

var (
	year = flag.Int("year", 2020, "The year to run")
)

func main() {
	t := time.Now()
	switch *year {
	case 2019:
		a2019.RunDay1("data/2019-1")
		runDay2()
		runDay3()
		runDay4()
		runDay5()
		runDay6()
		runDay7()
		runDay8()
		runDay9()
	case 2020:
		run2020day1()
		run2020day2()
		run2020day3()
		run2020day4()
		run2020day5()
		run2020day6()
		run2020day7()
		run2020day8()
	}
	fmt.Printf("Complete in %v", time.Now().Sub(t))
}

func loadProgram(file string) []int {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Cannot read data: %v", err)
	}

	program := []int{}
	elems := strings.Split(string(data), ",")
	for v, elem := range elems {
		i, err := strconv.ParseInt(elem, 10, 32)
		if err != nil {
			log.Fatalf("Bad conv at %v: %v", v, err)
		}
		program = append(program, int(i))
	}
	return program
}

func loadNums(file string) []int {
	lines := readLines(file)

	nums := []int{}
	for _, run := range lines {
		num, _ := strconv.ParseInt(string(run), 10, 32)
		nums = append(nums, int(num))
	}

	return nums
}

func loadListNums(file string) []int {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Cannot read data: %v", err)
	}

	nums := []int{}
	for _, run := range string(data) {
		num, _ := strconv.ParseInt(string(run), 10, 32)
		nums = append(nums, int(num))
	}

	return nums
}

func readLines(fileIn string) []string {
	file, err := os.Open(fileIn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func runDay2() {
	program := loadProgram("data/2019-2-1.txt")

	//Make called out adjustment
	program[1] = 12
	program[2] = 2
	fmt.Printf("Day2-1: %v\n", a2019.RunOpCode1(program))

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			program = loadProgram("data/2019-2-1.txt")

			program[1] = noun
			program[2] = verb

			result := a2019.RunOpCode1(program)
			if result == 19690720 {
				fmt.Printf("Day2-2: %v\n", 100*noun+verb)
			}
		}
	}
}

func runDay3() {
	file, err := os.Open("data/2019-3-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line1 := strings.Split(scanner.Text(), ",")
	scanner.Scan()
	line2 := strings.Split(scanner.Text(), ",")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day3-1: %v\n", a2019.ComputeDistance(line1, line2))
	fmt.Printf("Day3-2: %v\n", a2019.ComputeLen(line1, line2))
}

func runDay4() {
	count := 0
	count2 := 0
	for i := 134792; i <= 675810; i++ {
		if a2019.Valid(i) {
			count++
		}
		if a2019.TightValid(i) {
			count2++
		}
	}
	fmt.Printf("Day4-1: %v\n", count)
	fmt.Printf("Day4-2: %v\n", count2)
}

func runDay5() {
	program := loadProgram("data/2019-5-1.txt")
	result, _, _, _ := a2019.ProcessOpCode2(program, 0, []int{1}, 0, 0)
	fmt.Printf("Day5-1: %v\n", result)

	program = loadProgram("data/2019-5-1.txt")
	result, _, _, _ = a2019.ProcessOpCode2(program, 0, []int{5}, 0, 0)
	fmt.Printf("Day5-2: %v\n", result)
}

func runDay6() {
	lines := readLines("data/2019-6-1.txt")
	result, span := a2019.RunOrbits(lines)
	fmt.Printf("Day6-1: %v\n", result)
	fmt.Printf("Day6-2: %v\n", span)
}

func runDay7() {
	program := loadProgram("data/2019-7-1.txt")
	result := a2019.FindMax(program, []int{0, 1, 2, 3, 4})
	result2 := a2019.FindMax(program, []int{5, 6, 7, 8, 9})
	fmt.Printf("Day7-1: %v\n", result)
	fmt.Printf("Day7-2: %v\n", result2)
}

func runDay8() {
	nums := loadNums("data/2019-8-1.txt")
	result := a2019.FindLayer(nums, 25, 6)
	fmt.Printf("Day8-1: %v\n", result)
	a2019.PrintGraph(nums, 25, 6)
}

func runDay9() {
	program := loadProgram("data/2019-9-1.txt")
	output := a2019.RunFull(program)
	fmt.Printf("Day9-1: %v\n", output)
	output2 := a2019.RunFull2(program)
	fmt.Printf("Day9-2: %v\n", output2)
}
