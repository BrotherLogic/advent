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
	fmt.Printf("Day 7 [1]. %v\n", order)
}
