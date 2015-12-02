package main

import "io/ioutil"
import "strings"
import "fmt"

func ComputeFloor(str string) int {
	left := strings.Count(str, "(")
	right := strings.Count(str, ")")
	return left - right
}

func ComputeF1(str string) int {
	count := 0
	for i := 0; i < len(str) ; i++ {
		if str[i] == '(' {
			count++
		} else if str[i] == ')' {
			count--
		}

		if count == -1 {
			return i+1
		}
	}

	return -1
}

func dayone() {
	buf, err := ioutil.ReadFile("input")
	if err == nil {
		answer := ComputeFloor(string(buf))
		floor := ComputeF1(string(buf))
		fmt.Printf("Floor = %d\n", answer)
		fmt.Printf("First = %d\n", floor)
	} else {
		panic(err)
	}
}
