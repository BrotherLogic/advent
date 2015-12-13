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

func ProcJsonObject(dec *json.Decoder) (float64, bool) {
	count := 0.0
	seen_red := false
	for {
		t,err := dec.Token()
		if err == io.EOF || t == nil{
			break
		}

		if v,ok := t.(json.Delim); ok {
			if v.String() == "{" {
				ncount,seen := ProcJsonObject(dec)
				if !seen {
					count += ncount
				}
			} else if v.String() == "}" {
				return count,seen_red
			} else if v.String() == "[" {
				ncount,_ := ProcJsonObject(dec)
				count += ncount
			} else if v.String() == "]" {
				return count,seen_red
			}
		}

		if v,ok := t.(string); ok {
			if v == "red" {
				seen_red = true
			}
		}

		
		if v, ok := t.(float64); ok {
			count += v
		}

		
	}

	return count, seen_red
}

func ProcJsonStringNoRed(dec *json.Decoder) int {
	count, seen := ProcJsonObject(dec)
	if seen {
		return int(count)
	}
	return int(count)
}


func daytwelve() {
	file, e := os.Open("input-day12")
	file2, e := os.Open("input-day12")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	dec := json.NewDecoder(file)
	dec2 := json.NewDecoder(file2)
	fmt.Printf("Totat = %d\n", ProcJsonString(dec))
	fmt.Printf("Totat = %d\n", ProcJsonStringNoRed(dec2))
}
