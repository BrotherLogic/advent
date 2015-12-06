package main

import "bufio"
import "fmt"
import "log"
import "os"
import "regexp"
import "strconv"
import "strings"

func ToggleLights (str string, arr [][]bool) [][]bool {
	command := "on"
	if strings.Contains(str,"off") {
		command = "off"
	} else if strings.Contains(str,"toggle") {
		command = "toggle"
	}

	re := regexp.MustCompile(`(\d+),(\d+) through (\d+),(\d+)`)
	result := re.FindAllStringSubmatch(str,-1)

	x1,err := strconv.Atoi(result[0][1])
	y1,err := strconv.Atoi(result[0][2])
	x2,err := strconv.Atoi(result[0][3])
	y2,err := strconv.Atoi(result[0][4])

	if err != nil {
		panic("Halp")
	}
	
	for x := x1 ; x <= x2 ; x++ {
		for y := y1; y <= y2 ; y++ {
			switch(command) {
			case "on": arr[x][y] = true
			case "off": arr[x][y] = false
			case "toggle": arr[x][y] = !arr[x][y]
			}
		}
	}
	
	return arr
}

func ToggleBrightness (str string, arr [][]int) [][]int {
	command := "on"
	if strings.Contains(str,"off") {
		command = "off"
	} else if strings.Contains(str,"toggle") {
		command = "toggle"
	}

	re := regexp.MustCompile(`(\d+),(\d+) through (\d+),(\d+)`)
	result := re.FindAllStringSubmatch(str,-1)

	x1,err := strconv.Atoi(result[0][1])
	y1,err := strconv.Atoi(result[0][2])
	x2,err := strconv.Atoi(result[0][3])
	y2,err := strconv.Atoi(result[0][4])

	if err != nil {
		panic("Halp")
	}
	
	for x := x1 ; x <= x2 ; x++ {
		for y := y1; y <= y2 ; y++ {
			switch(command) {
			case "on": arr[x][y] = 0
			case "off": arr[x][y] = 0
			case "toggle": arr[x][y] = 0
			}
		}
	}
	
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

func MakeBrightBoard(size int) [][]int {
	arr := [][]int {}
	for i := 0 ; i < size ; i++ {
		row := make([]int, size)
		arr = append(arr, row)
	}
	return arr
}


func daysix() {
	lights := MakeBoard(1000)
	brights := MakeBrightBoard(1000)	

	file, err := os.Open("input-day6")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		lights = ToggleLights(text, lights)
		brights = ToggleBrightness(text,brights)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	count := 0
	bcount := 0
	for i := 0 ; i < len(lights) ; i++ {
		for j := 0 ; j < len(lights[i]) ; j++ {
			if (lights[i][j]) {
				count++
			}
			bcount += brights[i][j]
		}
	}

	fmt.Printf("Lights = %d\n", count)
	fmt.Printf("Brights = %d\n", bcount)
}
