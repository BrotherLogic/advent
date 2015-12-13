package main

import "fmt"
import "encoding/json"
import "io"
import "os"

func ProcJsonString(dec *json.Decoder) int {
	count := 0.0
	for {
		t, err := dec.Token()
		
		if err == io.EOF {
			break
		}

		if v, ok := t.(float64); ok {
			count += v
		}
		
		if t == nil {
			break
		}
	}

	return int(count)
}

func ProcJsonStringNoRed(dec *json.Decoder) int {
	count := 0.0
	for {
		t, err := dec.Token()
		
		if err == io.EOF {
			break
		}

		if v, ok := t.(float64); ok {
			count += v
		}
		
		if t == nil {
			break
		}
	}

	return int(count)
}


func daytwelve() {
	file, e := os.Open("input-day12")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	dec := json.NewDecoder(file)
	fmt.Printf("Totat = %d\n", ProcJsonString(dec))
	fmt.Printf("Totat = %d\n", ProcJsonStringNoRed(dec))
}
