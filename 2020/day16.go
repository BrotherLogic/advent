package a2020

import (
	"log"
	"strconv"
	"strings"
)

type nrange struct {
	minv int
	maxv int
}

func (n *nrange) inRange(num int64) bool {
	return n.minv <= int(num) && n.maxv >= int(num)
}

func rconvert(str string) nrange {
	parts := strings.Split(str, "-")
	c1, _ := strconv.ParseInt(parts[0], 10, 32)
	c2, _ := strconv.ParseInt(parts[1], 10, 32)
	return nrange{minv: int(c1), maxv: int(c2)}
}

func buildRep(lines []string) map[string][]nrange {
	nmap := make(map[string][]nrange)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) == 2 && len(parts[1]) > 0 {
			ident := parts[0]
			fields := strings.Fields(parts[1])
			nmap[ident] = []nrange{rconvert(fields[0]), rconvert(fields[2])}
		}
	}
	return nmap
}

func buildColumnMap(lines []string) map[string]int {
	rep := buildRep(lines)
	pos := make(map[string][]int)

	for key := range rep {
		arr := []int{}
		for i := 0; i < len(rep); i++ {
			arr = append(arr, i)
		}
		pos[key] = arr
	}

	_, goods := FindTicket(lines)
	for _, linenum := range goods {
		line := lines[linenum]
		nums := strings.Split(line, ",")
		for i := range nums {
			num, _ := strconv.ParseInt(nums[i], 10, 32)
			for key, vals := range pos {
				nvals := []int{}
				for _, val := range vals {
					found := i != val
					for _, r := range rep[key] {
						if r.inRange(num) {
							found = true
						}
					}
					if found {
						nvals = append(nvals, val)
					}
				}
				pos[key] = nvals
			}
		}
	}

	return resolveMap(pos)
}

func resolveMap(m map[string][]int) map[string]int {
	r := make(map[string]int)

	valsdone := []int{}
	count := 0
	for len(r) != len(m) {
		for key, val := range m {
			if len(val) == 1 {
				r[key] = val[0]
				valsdone = append(valsdone, val[0])
			}
		}

		for key, val := range m {
			newval := []int{}
			for _, v := range val {
				found := false
				for _, v2 := range valsdone {
					if v == v2 {
						found = true
					}
				}
				if !found {
					newval = append(newval, v)
				}
			}
			m[key] = newval
		}

		count++
		if count > 1000 {
			log.Fatalf("TOO MUCH WORK!")
		}
	}

	return r
}

//BuildTicket builds out the ticket
func BuildTicket(lines []string, pr string) int {
	m := buildColumnMap(lines)

	mine := ""
	fmine := false
	for _, line := range lines {
		if fmine {
			mine = line
			fmine = false
		}
		if line == "your ticket:" {
			fmine = true
		}
	}
	elems := strings.Split(mine, ",")

	sumv := 1
	for key, row := range m {
		if strings.HasPrefix(key, pr) {
			elem := elems[row]
			num, _ := strconv.ParseInt(elem, 10, 32)
			sumv *= int(num)
		}
	}

	return sumv
}

//FindTicket finds the ticket
func FindTicket(lines []string) (int, []int) {
	scansum := int64(0)
	goods := []int{}
	rep := buildRep(lines)

	checking := false
	for i, line := range lines {
		invalid := false
		if checking {
			nums := strings.Split(line, ",")
			for _, num := range nums {
				nv, _ := strconv.ParseInt(num, 10, 32)
				found := false
				for _, ra := range rep {
					for _, r := range ra {
						if r.inRange(nv) {
							found = true
							break
						}
					}
				}

				if !found {
					scansum += nv
					invalid = true
				}
			}

			if !invalid {
				goods = append(goods, i)
			}
		}

		if !checking && strings.HasPrefix(line, "nearby tickets") {
			checking = true
		}
	}

	return int(scansum), goods
}
