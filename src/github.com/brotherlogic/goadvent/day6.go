package main

import "bufio"
import "fmt"
import "log"
import "os"

func ToggleLights (str string, arr [][]bool) [][]bool {
	return arr
}

func MakeBoard(size int) [][]bool {
	arr := [][]bool {}
	for i := 0 ; i < size ; i++ {
		row := make([]bool, size)
		arr = append(arr, row)
	}
	return arr
}

func daysix() {
	lights := MakeBoard(1000)
	

	file, err := os.Open("input-day6")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		lights = ToggleLights(text, lights)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	count := 0
	for i := 0 ; i < len(lights) ; i++ {
		for j := 0 ; j < len(lights[i]) ; j++ {
			if (lights[i][j]) {
				count++
			}
		}
	}

	fmt.Printf("Lights = %d\n", count)
}
