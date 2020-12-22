package a2020

import (
	"log"
	"strconv"
	"strings"
)

type melem struct {
	number   int
	operator string
}

func nconvert(str string) melem {
	if str == "*" || str == "+" {
		return melem{operator: str}
	}

	n, _ := strconv.ParseInt(str, 10, 32)
	return melem{number: int(n)}
}

func removeSpaces(line string) []melem {
	fields := strings.Fields(line)
	elems := []melem{}
	for _, f := range fields {
		lcount := strings.Count(f, "(")
		rcount := strings.Count(f, ")")

		for i := 0; i < lcount; i++ {
			elems = append(elems, melem{operator: "("})
		}
		elems = append(elems, nconvert(f[lcount:len(f)-rcount]))

		for i := 0; i < rcount; i++ {
			elems = append(elems, melem{operator: ")"})
		}

	}

	return elems
}

func convertToRP(line []melem, prec map[string]int) []melem {
	o := []melem{}
	s := []melem{}

	for _, token := range line {
		if token.number > 0 {
			o = append(o, token)
		}
		if token.operator == "*" || token.operator == "+" {
			if len(s) > 0 {
				if s[0].operator == "*" || s[0].operator == "+" {
					if prec[s[0].operator] >= prec[token.operator] {
						o = append(o, s[0])
						s = s[1:]
					}
				}
			}
			s = append([]melem{token}, s...)
		}

		if token.operator == "(" {
			s = append([]melem{token}, s...)
		}

		if token.operator == ")" {
			for s[0].operator != "(" {
				o = append(o, s[0])
				s = s[1:]
			}
			s = s[1:]
		}
	}

	for _, t := range s {
		o = append(o, t)
	}

	return o
}

func computeCV(cv []melem) int {
	count := 0
	for len(cv) > 1 {
		for i := 0; i < len(cv)-2; i++ {
			if cv[i].number > 0 && cv[i+1].number > 0 && cv[i+2].operator != "" {
				ncv := cv[:i]

				sol := cv[i].number + cv[i+1].number
				if cv[i+2].operator == "*" {
					sol = cv[i].number * cv[i+1].number
				}

				ncv = append(ncv, melem{number: sol})

				ncv = append(ncv, cv[i+3:]...)
				cv = ncv
				break
			}
		}
		count++
		if count > 100 {
			log.Fatalf("TOO MANY RUNS")
		}

	}
	return cv[0].number
}

func runSingleSum(line string, prec map[string]int) int {
	nsp := removeSpaces(line)
	//log.Printf("HERE: %v", nsp)
	cv := convertToRP(nsp, prec)
	//log.Printf("%v -> %v", line, cv)
	return computeCV(cv)
	//return 0
}

//RunSum runs the full sum
func RunSum(lines []string, prec map[string]int) int {
	sumv := 0
	for _, line := range lines {
		sumv += runSingleSum(line, prec)
	}
	return sumv
}
