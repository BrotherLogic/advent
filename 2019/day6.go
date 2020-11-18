package a2019

import "strings"

func computeOrbits(mapper map[string]string, key string) int {
	nkey := key
	count := 0
	for nkey != "COM" {
		count++
		nkey = mapper[nkey]
	}
	return count
}

//RunOrbits runs the orbits
func RunOrbits(orbits []string) (int, int) {
	mapper := make(map[string]string)

	for _, entry := range orbits {
		parts := strings.Split(entry, ")")
		mapper[parts[1]] = parts[0]
	}

	orbitCount := 0
	for key := range mapper {
		orbitCount += computeOrbits(mapper, key)
	}

	youpath := []string{}
	pointer := "YOU"
	for pointer != "COM" {
		youpath = append(youpath, pointer)
		pointer = mapper[pointer]
	}

	pointer = "SAN"
	counter := 0
	for pointer != "COM" {
		for i, pp := range youpath {
			if pp == mapper[pointer] {
				return orbitCount, i + counter - 1
			}
		}
		counter++
		pointer = mapper[pointer]
	}

	return orbitCount, 0
}
