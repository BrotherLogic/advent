package main

import "io/ioutil"
import "fmt"
import "strconv"

func ComputeNumberOfRoboHouses(str string) int {
	return 0
}

func ComputeNumberOfHouses(str string) int {
	x := 0
	y := 0

	var m map[string]int
	m = make(map[string]int)

	m[strconv.Itoa(x) + "|" + strconv.Itoa(y)] = 1

	for i := 0 ; i < len(str) ; i++ {
		switch str[i] {
		case '<':
			x--
		case '>':
			x++
		case '^':
			y++
		case 'v':
			y--
		default:
			panic("Unknown string")
		}

		m[strconv.Itoa(x) + "|" + strconv.Itoa(y)] = 1
	}
	
	return len(m)
}

func daythree() {
	buf, err := ioutil.ReadFile("input-day3")
	if err == nil {
		answer := ComputeNumberOfHouses(string(buf))
		roboanswer := ComputeNumberOfRoboHouses(string(buf))
		fmt.Printf("Number of Houses = %d\n", answer)
		fmt.Printf("Number of RoboHouses = %d\n", roboanswer)
	} else {
		panic(err)
	}
}
