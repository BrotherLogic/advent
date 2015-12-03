package main

import "io/ioutil"
import "fmt"

func ComputeNumberOfHouses(str string) int {
	return 0
}

func daythree() {
	buf, err := ioutil.ReadFile("input-day3")
	if err == nil {
		answer := ComputeNumberOfHouses(string(buf))
		fmt.Printf("Number of Houses = %d\n", answer)
	} else {
		panic(err)
	}
}
