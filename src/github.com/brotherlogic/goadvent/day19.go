package main

import "bufio"
import "fmt"
import "os"
import "strings"

func ReverseMap(mapper map[string][]string) map[string][]string {
	var nmap map[string][]string
	nmap = make(map[string][]string)

	for key, val := range mapper {
		for i := 0 ; i < len(val) ; i++ {
			if _, ok := nmap[val[i]]; ok {
				nmap[val[i]] = append(nmap[val[i]], key)
			} else {
				nmap[val[i]] = make([]string,0)
				nmap[val[i]] = append(nmap[val[i]], key)
			}
		}
	}
	return nmap
}

func ReverseMapping(mapper map[string][]string, target string, steps int) int {
	best := 9999
	
	if target == "e" {
		return steps
	}

	if strings.Contains(target,"e") {
		return best
	}
	
	for key, val := range mapper {
		strs := strings.Split(target,key)
		for i:= 0 ; i < len(strs)-1 ; i++ {
			str := ""
			nfirst := true
			for j := 0 ; j <= i ; j++ {
				if nfirst {
					str += strs[j]
					nfirst = false
				} else {
					str += key + strs[j]
				}
			}
			for k := 0 ; k < len(val) ; k++ {
				nstr := str + val[k]
				first := true
				for l := i+1 ; l < len(strs) ; l++ {
					if first {
						nstr += strs[l]
						first = false
					} else {
						nstr += key + strs[l]
					}
				}
				best = Min(best,ReverseMapping(mapper, nstr, steps +1))
			}
		}
	}
	
	return best
}

func Add(mapper map[string]int, str string) map[string]int {
	if _, ok := mapper[str]; ok {
		mapper[str] += 1
	} else {
		mapper[str] = 1
	}

	return mapper
}

func BuildMapping(mapper map[string][]string, pointer int, input string) map[string]int {
	
	var found map[string]int
	found = make(map[string]int)
	if pointer >= len(input) {
		return found
	}
	
	char := ""
	
	if pointer < len(input)-1 && strings.ToLower(input[pointer+1:pointer+2]) == input[pointer+1:pointer+2] {
		char = input[pointer:pointer+2]
	} else {
		char = input[pointer:pointer+1]
	}

	if val, ok := mapper[char]; ok {
		for i := 0 ; i < len(val) ; i++ {
			nstring := input[0:pointer] + val[i] + input[pointer+len(char):len(input)]		
			Add(found,nstring)
		}
	}
	
	nmap := BuildMapping(mapper, pointer+len(char), input)
	for k, v := range nmap {
		found[k] += v
	}
	
	
	return found
}

func daynineteen() {
	file,_ := os.Open("input-day19")
	scanner := bufio.NewScanner(file)
	var mapper map[string][]string
	mapper = make(map[string][]string)
	
	for scanner.Scan() {
		text := scanner.Text()
		index := strings.Index(text," => ")
		if index >= 0 {
			t1 := text[0:index]
			t2 := text[index+4:len(text)]

			if _, ok := mapper[t1]; ok {
				mapper[t1] = append(mapper[t1],t2)
			} else {
				mapper[t1] = make([]string,0)
				mapper[t1] = append(mapper[t1],t2)
			}
		} else if len(text) > 2 {
			fmt.Printf("Answer = %v\n", len(BuildMapping(mapper, 0, text)))
			fmt.Printf("Steps = %v\n", ReverseMapping(mapper, text, 0))
		}
	}
}
