package a2020

import (
	"strconv"
	"strings"
)

func convertToPassport(lines []string) []map[string]string {
	res := make([]map[string]string, 0)

	mapper := make(map[string]string)
	for _, line := range lines {
		if line == "" {
			res = append(res, mapper)
			mapper = make(map[string]string)
		} else {
			fields := strings.Fields(line)
			for _, field := range fields {
				split := strings.Split(field, ":")
				mapper[split[0]] = split[1]
			}
		}
	}
	res = append(res, mapper)
	return res
}

func isValid1(m map[string]string) bool {
	for _, need := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if m[need] == "" {
			return false
		}
	}

	return true
}

func isValid2(m map[string]string) bool {
	for _, need := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if m[need] == "" {
			return false
		}
		if need == "byr" {
			num, err := strconv.ParseInt(m[need], 10, 32)
			if err != nil || num < 1920 || num > 2002 {
				return false
			}
		}
		if need == "iyr" {
			num, err := strconv.ParseInt(m[need], 10, 32)
			if err != nil || num < 2010 || num > 2020 {
				return false
			}
		}
		if need == "eyr" {
			num, err := strconv.ParseInt(m[need], 10, 32)
			if err != nil || num < 2020 || num > 2030 {
				return false
			}
		}
		if need == "hgt" {
			end := string(m[need][len(m[need])-2:])
			num, err := strconv.ParseInt(m[need][:len(m[need])-2], 10, 32)
			if end != "cm" && end != "in" {
				return false
			}
			if end == "cm" {
				if err != nil || num < 150 || num > 193 {
					return false
				}
			}
			if end == "in" {
				if err != nil || num < 59 || num > 76 {
					return false
				}
			}
		}
		if need == "hcl" {
			if m[need][0] != '#' {
				return false
			}
			for _, r := range m[need][1:] {
				if r != '1' && r != '2' && r != '3' && r != '4' && r != '5' && r != '6' &&
					r != '7' && r != '8' && r != '9' && r != '0' && r != 'a' && r != 'b' && r != 'c' && r != 'd' && r != 'e' && r != 'f' {
					return false
				}
			}
		}
		if need == "ecl" {
			if m[need] != "amb" &&
				m[need] != "blu" &&
				m[need] != "brn" &&
				m[need] != "gry" &&
				m[need] != "grn" &&
				m[need] != "hzl" &&
				m[need] != "oth" {
				return false
			}
		}
		if need == "pid" {
			if len(m[need]) != 9 {
				return false
			}
			for _, r := range m[need] {
				if r != '1' && r != '2' && r != '3' && r != '4' && r != '5' && r != '6' &&
					r != '7' && r != '8' && r != '9' && r != '0' {
					return false
				}
			}
		}
	}

	return true
}

//CountValid1 the number of valid passports
func CountValid1(lines []string) int {
	count := 0
	for _, pass := range convertToPassport(lines) {
		if isValid1(pass) {
			count++
		}
	}
	return count
}

//CountValid2 the number of valid passports
func CountValid2(lines []string) int {
	count := 0
	for _, pass := range convertToPassport(lines) {
		if isValid2(pass) {
			count++
		}
	}
	return count
}
