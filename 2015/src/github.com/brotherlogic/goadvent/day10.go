package main

import "bytes"
import "fmt"
import "strconv"

func LookAndSay(str string) string {
	var newstr bytes.Buffer
	curr_char := str[0]
	curr_count := 1

	for i := 1; i < len(str); i++ {
		if str[i] != curr_char {
			newstr.WriteString(strconv.Itoa(curr_count) + string(curr_char))
			curr_char = str[i]
			curr_count = 1
		} else {
			curr_count++
		}
	}

	newstr.WriteString(strconv.Itoa(curr_count) + string(curr_char))
	return newstr.String()
}

func LookApply(str string, count int) string {
	for i := 0 ; i < count ; i++ {
		str  = LookAndSay(str)
	}
	return str
}

func dayten() {
	init := "1321131112"	
	fmt.Printf("Part 1 = %d\n", len(LookApply(init, 40)))
	fmt.Printf("Part 2 = %d\n", len(LookApply(init, 50)))
}
