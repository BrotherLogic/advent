package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func shortest(in string) int {
	res := make(map[rune]int)

	for i := 0; i < len(in); i++ {
		lcr := unicode.ToLower(rune(in[i]))
		if _, ok := res[lcr]; !ok {
			newString := ""
			for j := 0; j < len(in); j++ {
				if unicode.ToLower(rune(in[j])) != lcr {
					newString += string(in[j])
				}
			}

			res[lcr] = len(collapse(newString))
		}
	}

	maxV := len(in)
	for _, val := range res {
		if val < maxV {
			maxV = val
		}
	}

	return maxV
}

func collapse(in string) string {
	collapsed := in
	i := 0
	for i < len(collapsed)-1 {
		if i < 0 {
			i = 0
		}

		if collapsed[i] != collapsed[i+1] && unicode.ToUpper(rune(collapsed[i])) == unicode.ToUpper(rune(collapsed[i+1])) {
			collapsed = collapsed[0:i] + collapsed[i+2:len(collapsed)]
			i--
		} else {
			i++
		}
	}

	return collapsed
}

func solveDay5(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	text := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 5[1]. %v\n", len(collapse(text)))
	fmt.Printf("Day 5[2]. %v\n", shortest(text))
}
