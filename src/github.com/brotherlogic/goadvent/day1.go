package main

import "io/ioutil"
import "strings"
import "fmt"

func ComputeFloor(str string) int {
	left := strings.Count(str, "(")
	right := strings.Count(str, ")")
	fmt.Printf("Found %d,%d\n", left, right)
	return left - right
}

func main() {
	buf, err := ioutil.ReadFile("input")
	if err == nil {
		answer := ComputeFloor(string(buf))
		fmt.Printf("Floor = %d\n", answer)
	} else {
		panic(err)
	}
}
