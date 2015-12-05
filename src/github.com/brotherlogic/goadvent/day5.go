package main

import "bufio"
import "fmt"
import "log"
import "os"

func IsNice(str string) bool {
	return true
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
