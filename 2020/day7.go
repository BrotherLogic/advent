package a2020

import (
	"fmt"
	"strconv"
	"strings"
)

type child struct {
	count int
	name  string
}

func splitLine(line string) (string, []child) {
	fs := strings.Fields(line)
	others := []child{}
	if fs[4] != "no" {
		for i := 5; i < len(fs); i += 4 {
			c, _ := strconv.ParseInt(fs[i-1], 10, 32)
			others = append(others, child{name: fmt.Sprintf("%v %v", fs[i], fs[i+1]), count: int(c)})
		}
	}
	return fmt.Sprintf("%v %v", fs[0], fs[1]), others
}

func find(m map[string][]child, head, ty string) bool {
	if head == ty {
		return true
	}

	for _, res := range m[head] {
		if find(m, res.name, ty) {
			return true
		}
	}

	return false
}

//Derive the bag count
func Derive(lines []string, ty string) int {
	m := make(map[string][]child)
	for _, line := range lines {
		head, children := splitLine(line)
		m[head] = children
	}

	count := 0
	for head := range m {
		if find(m, head, ty) {
			count++
		}
	}

	// Less one since a bag cannot contain itself
	return count - 1
}

func doCount(m map[string][]child, ty string) int {
	count := 1

	if len(m[ty]) == 0 {
		return 1
	}

	for _, others := range m[ty] {
		count += others.count * doCount(m, others.name)
	}
	return count
}

// Delve into the bag count
func Delve(lines []string, ty string) int {
	m := make(map[string][]child)
	for _, line := range lines {
		head, children := splitLine(line)
		m[head] = children
	}

	return doCount(m, ty) - 1
}
