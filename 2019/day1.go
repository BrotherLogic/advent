package a2019

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func computeSuperFuel(mass int) int {
	sum := -mass
	for mass > 0 {
		sum += mass
		mass = computeFuel(mass)
	}
	return sum
}

func computeFuel(mass int) int {
	return mass/3 - 2
}

func RunDay1(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	superSum := 0
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += computeFuel(i)
		superSum += computeSuperFuel(i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 1, Part 1 = %v\n", sum)
	fmt.Printf("Day 1, Part 2 = %v\n", superSum)
}
