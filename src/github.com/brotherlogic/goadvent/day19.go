package main

import "bufio"
import "fmt"
import "os"
import "math/rand"
import "sort"
import "strings"
import "time"

func ShuffleStrings(slc []string) {
	rand.Seed(time.Now().UnixNano())
	N := len(slc)
	for i := 0; i < N; i++ {
		// choose index uniformly in [i, N-1]
		r := i + rand.Intn(N-i)
		slc[r], slc[i] = slc[i], slc[r]
	}
}

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

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func ReverseMapping2(mapper map[string][]string, target string, steps int) int {
	
	best := 99999
	
	var keys []string
	for k := range mapper {
		keys = append(keys,k)
	}

	if target == "e" {
		return steps
	}
	
	//sort.Sort(ByLength(keys))
	ShuffleStrings(keys)

//	fmt.Printf("Starting %v\n", keys)
	ntarget := target
	for ntarget != "e" {
		breaker := false
		for _, key := range keys {
			if strings.Contains(ntarget, key) {
//				fmt.Printf("HERE = %v with %v => %v, %v\n",ntarget, key, mapper[key], steps)
				breaker = true
				ntarget = strings.Replace(ntarget,key,mapper[key][0], 1)
				steps++
			}
		}

		if !breaker || steps > 1000 {
//			fmt.Printf("Could not find %v,%v,%v\n", breaker, steps, ntarget)
			return 5555
		}
	}

	return steps
		
	for _, key := range keys {
		if strings.Contains(target, key) {
			arr := mapper[key]
			sort.Sort(ByLength(arr))

			repl := strings.Replace(target, key, arr[0], 1)
			best = Min(best,ReverseMapping2(mapper, repl, steps+1))
		}
	}

	return best
}

func ReverseMapping(mapper map[string][]string, target string, steps int) int {

	best := 9999
	
	if target == "e" {
		return steps
	}

	if strings.Contains(target,"e") {
		return best
	}

	var keys []string
	for k := range mapper {
		keys = append(keys,k)
	}

	sort.Sort(ByLength(keys))
	for _, key := range keys {
		val := mapper[key]
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
				
				if len(nstr) <= len(target) {
					best = Min(best,ReverseMapping(mapper, nstr, steps +1))
				} 
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
			for i := 0 ; i < 20000 ; i++ {
				found := ReverseMapping2(ReverseMap(mapper), text, 0)
				if found < 5000 {
					fmt.Printf("Steps = %v\n", found)
				}
			}
		}
	}
}
