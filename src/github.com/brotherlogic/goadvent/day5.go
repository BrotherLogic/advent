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

func CountMaxNonOverlapping(str string) int {
	counts := make(map[string]int)
	for i := 0 ; i < len(str) -1; i++ {
		pair := string(str[i:i+2])
		safe := true
		if str[i] == str[i+1] {
			if i > 0 && str[i] == str[i-1] {
				safe = false
			}
			if i < len(str)-2 && str[i] == str[i+2] {
				safe = false
			}
		}

		if safe {
			counts[pair]++
		}
	}

	max_v := 0
	for _, value := range counts {
		max_v = Max(value,max_v)
	}
	return max_v
}

func RepeatWithMiddle(str string) bool {
	for i := 0 ; i < len(str)-2 ; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
}

func IsNiceAlso(str string) bool {
	fmt.Printf("%q = %d, %t\n",str,CountMaxNonOverlapping(str),RepeatWithMiddle(str))
	return CountMaxNonOverlapping(str) >= 2 && RepeatWithMiddle(str)
}

func IsNice(str string) bool {
	bad_str := []string{"ab","cd","pq","xy"}
	for i := 0 ; i < len(bad_str) ; i++ {
		if strings.Contains(str,bad_str[i]) {
			return false
		}
	}
	
	
	return CountVowels(str) >= 3 && CountRow(str) >= 2
}

func dayfive() {
	file, err := os.Open("input-day5")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	nice := 0
	also_nice := 0
	for scanner.Scan() {
		text := scanner.Text()
		if IsNice(text) {
			nice++
		}
		if IsNiceAlso(text) {
			also_nice++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Nice = %d\n",nice)
	fmt.Printf("Also Nice = %d\n",also_nice)
}
