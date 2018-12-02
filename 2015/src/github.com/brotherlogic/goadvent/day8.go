package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"

func Convert(str string) int {
	length_un := len(str)
	sc,_ := strconv.Unquote(str)
	length_dn := len(sc)
	
	return length_un - length_dn
}

func Convert2(str string) int {
	length_un := len(str)
	length_dn := len(strconv.Quote(str))

	return length_dn-length_un
}

func dayeight() {
	
	file, err := os.Open("input-day8")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	count := 0
	count2 := 0
	for scanner.Scan() {
		text := scanner.Text()
		count += Convert(text)
		count2 += Convert2(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Answer = %d, %d\n",count, count2)
}
