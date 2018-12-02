package main

import "fmt"

func Next(in []byte) []byte {
	for i := len(in)-1; i >= 0; i-- {
		in[i]++
		if in[i] == 'z'+1 {
			in[i] = 'a'
		} else {
			break
		}
	}
	return in
}

func FirstReq(in []byte) bool {
	inc_count := 1
	inc_previous := in[0]

	for i := 1; i < len(in); i++ {
		if in[i] == inc_previous+1 {
			inc_previous++
			inc_count++

			if inc_count == 3 {
				return true
			}
		} else {
			inc_count = 1
			inc_previous = in[i]
		}
	}

	return false
}

func SecondReq(in []byte) bool {
	for i := 0 ; i < len(in); i++ {
		if in[i] == 'i' || in[i] == 'o' || in[i] == 'l' {
			return false
		}
	}
	return true
}

func ThirdReq(in []byte) bool {
	pair_count := 0
	for i := 1; i < len(in); i++ {
		if in[i] == in[i-1] {
			pair_count++
			i++
		}
	}

	return pair_count >= 2
}

func NewPassword(str string) string {
	next := Next([]byte(str))
	for !(FirstReq(next) && SecondReq(next) && ThirdReq(next)) {
		next = Next(next)
	}
	return string(next)
}

func dayeleven() {
	init := "hxbxwxba"	
	fmt.Printf("Part 1 = %q\n", NewPassword(init))
	fmt.Printf("Part 2 = %q\n", NewPassword(NewPassword(init)))
}
