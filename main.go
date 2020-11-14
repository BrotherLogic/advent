package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	a2019 "github.com/brotherlogic/advent/2019"
)

func main() {
	a2019.RunDay1("data/2019-1")
	runDay2()
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
