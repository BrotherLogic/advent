package main

import "io/ioutil"
import "fmt"
import "strconv"

func ComputeNumberOfRoboHouses(str string) int {
	hx := 0
	hy := 0
	rx := 0
	ry := 0
	
	var m map[string]int
	m = make(map[string]int)

	m[strconv.Itoa(hx) + "|" + strconv.Itoa(hy)] = 1

	for i := 0 ; i < len(str) ; i++ {
		switch str[i] {
		case '<':
			if i % 2 == 0 {
				hx--
			} else {
				rx--
			}
		case '>':
			if i % 2 == 0 {
				hx++
			} else {
				rx++
			}
		case '^':
			if i % 2 == 0 {
				hy++
			} else {
				ry++
			}

		case 'v':
			if i % 2 == 0 {
				hy--
			} else {
				ry--
			}

		default:
			panic("Unknown string")
		}

		m[strconv.Itoa(hx) + "|" + strconv.Itoa(hy)] = 1
		m[strconv.Itoa(rx) + "|" + strconv.Itoa(ry)] = 1
	}
	
	return len(m)
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
