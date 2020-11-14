package a2019

import (
	"strconv"
)

func digit(num, index int) int {
	str := strconv.Itoa(num)
	val, _ := strconv.ParseInt(string(str[index]), 10, 32)
	return int(val)
}

//Valid password?
func Valid(num int) bool {

	//Adjacents are the same
	foundr := false
	founds := false
	for i := 0; i < 5; i++ {
		if digit(num, i) == digit(num, i+1) {
			foundr = true
		}
		if digit(num, i) > digit(num, i+1) {
			founds = true
		}
	}
	if !foundr || founds {
		return false
	}

	return true
}

//TightValid password?
func TightValid(num int) bool {

	//Adjacents are the same
	foundr := false
	founds := false
	for i := 0; i < 5; i++ {
		if digit(num, i) == digit(num, i+1) && (i == 4 || digit(num, i) != digit(num, i+2)) && (i == 0 || digit(num, i) != digit(num, i-1)) {
			foundr = true
		}

		if digit(num, i) > digit(num, i+1) {
			founds = true
		}
	}
	if !foundr || founds {
		return false
	}

	return true
}
