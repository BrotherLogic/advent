package main

import "fmt"
import "strconv"

func LookAndSay(str string) string {
	newstr := ""
	curr_char := str[0]
	curr_count := 1

	for i := 1; i < len(str); i++ {
		if str[i] != curr_char {
			newstr += strconv.Itoa(curr_count) + string(curr_char)
			curr_char = str[i]
			curr_count = 1
		} else {
			curr_count++
		}
	}

	newstr += strconv.Itoa(curr_count) + string(curr_char)
	return newstr
}

func LookApply(str string, count int) string {
	for i := 0 ; i < count ; i++ {
		str  = LookAndSay(str)
	}
	return str
}

func dayten() {
	init := "1321131112"	
	fmt.Printf("Part 1 = %q\n", LookApply(init, 40))
}
