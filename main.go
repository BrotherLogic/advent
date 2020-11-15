package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	a2019 "github.com/brotherlogic/advent/2019"
)

func main() {
	a2019.RunDay1("data/2019-1")
	runDay2()
	runDay3()
	runDay4()
	runDay5()
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
	result := a2019.ProcessOpCode2(program, 1)
	fmt.Printf("Day5-1: %v\n", result)

	program = loadProgram("data/2019-5-1.txt")
	result = a2019.ProcessOpCode2(program, 5)
	fmt.Printf("Day5-2: %v\n", result)
}
