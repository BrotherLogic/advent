package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strings"

func CountVowels(str string) int {
	count := 0
	for i := 0 ; i < len(str) ; i++ {
		switch str[i] {
		case 'a','e','i','o','u':
			count++
		}
	}
	return count
}

func CountRow(str string) int {
	row_count := 1
	best := 0
	current := str[0]

	for i := 1 ; i < len(str) ; i++ {
		if str[i] == current {
			row_count++
			best = Max(best,row_count)
		} else {
			row_count = 1
		}
		current = str[i]
	}
	
	return best
}

func IsNice(str string) bool {
	bad_str := []string{"ab","cd","pq","xy"}
	for i := 0 ; i < len(bad_str) ; i++ {
		if strings.Contains(str,bad_str[i]) {
			fmt.Printf("%q contains bad\n",str)
			return false
		}
	}
	
	
	return CountVowels(str) >= 3 && CountRow(str) >= 2
}

func dayfive() {
	file, err := os.Open("input-day2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nice := 0
	for scanner.Scan() {
		text := scanner.Text()
		if IsNice(text) {
			nice++;
		}
	}

	fmt.Printf("Nice = %d\n",nice)
}
