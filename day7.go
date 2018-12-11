package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type worker struct {
	workingOn    rune
	timeFinished int
}

func resolveTime(order []string, maxv int, workers int, addition int) int {
	answer := ""
	timeNow := 0
	workArr := make([]worker, workers)
	for i := range workArr {
		workArr[i] = worker{workingOn: ' ', timeFinished: 0}
	}
	done := make(map[rune]bool)
	inprof := make(map[rune]bool)
	for len(answer) != maxv {
		for i := range workArr {
			if workArr[i].workingOn != ' ' && workArr[i].timeFinished == timeNow {
				done[workArr[i].workingOn] = true
				answer += string(workArr[i].workingOn)

				workArr[i].workingOn = ' '
			}

			for ct, char := range alpha[:maxv] {
				if !done[char] && !inprof[char] {
					cando := true
					for _, ord := range order {
						if rune(ord[36]) == char && !done[rune(ord[5])] {
							cando = false
							break
						}
					}

					if cando {
						for i := range workArr {
							if workArr[i].workingOn == ' ' {
								log.Printf("%v worker %v working on %v", timeNow, i, string(char))
								inprof[char] = true
								workArr[i].workingOn = char
								workArr[i].timeFinished = timeNow + ct + 1 + addition
								break
							}
						}
					}
				}
			}
		}
		timeNow++
	}
	return timeNow - 1
}

func resolveOrder(order []string, maxv int) string {
	answer := ""

	done := make(map[rune]bool)
	for len(answer) != maxv {
		for _, char := range alpha[:maxv] {
			if !done[char] {
				cando := true
				for _, ord := range order {
					if rune(ord[36]) == char && !done[rune(ord[5])] {
						cando = false
						break
					}
				}

				if cando {
					answer += string(char)
					done[char] = true
					break
				}
			}
		}
	}

	return answer
}

func solveDay7(file string) {
	arr := []string{}
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	order := resolveOrder(arr, 26)
	timeV := resolveTime(arr, 26, 5, 60)
	fmt.Printf("Day 7 [1]. %v\n", order)
	fmt.Printf("Day 7 [2]. %v\n", timeV)
}
