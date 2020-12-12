package a2020

import "log"

func doCopy(lines []string) []string {
	cp := make([]string, len(lines))
	copy(cp, lines)
	return cp
}

func runAlgorithm(seats []string) []string {
	nseats := doCopy(seats)
	for y := 0; y < len(seats); y++ {
		for x := 0; x < len(seats[y]); x++ {

			occupied := 0
			for ry := maxint(0, y-1); ry <= minint(y+1, len(seats)-1); ry++ {
				for rx := maxint(0, x-1); rx <= minint(x+1, len(seats[0])-1); rx++ {
					if ry != y || rx != x {
						if seats[ry][rx] == '#' {
							occupied++
						}
					}
				}
			}

			if seats[y][x] == 'L' && occupied == 0 {
				nseats[y] = nseats[y][:x] + "#" + nseats[y][x+1:]
			}
			if seats[y][x] == '#' && occupied >= 4 {
				nseats[y] = nseats[y][:x] + "L" + nseats[y][x+1:]
			}

		}
	}

	return nseats
}

type dir struct {
	x, y int
}

func runAlgorithm2(seats []string) []string {
	nseats := doCopy(seats)

	dirs := []dir{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x != 0 || y != 0 {
				dirs = append(dirs, dir{x: x, y: y})
			}
		}
	}

	for y := 0; y < len(seats); y++ {
		for x := 0; x < len(seats[y]); x++ {

			occupied := 0
			for _, d := range dirs {
				found := false
				for mult := 1; mult < len(seats); mult++ {
					nx, ny := x+d.x*mult, y+d.y*mult
					if nx >= 0 && nx < len(seats[y]) && ny >= 0 && ny < len(seats) {
						if seats[ny][nx] == '#' {
							found = true
							break
						} else if seats[ny][nx] == 'L' {
							break
						}
					} else {
						break
					}
				}

				if found {
					occupied++
				}
			}

			if seats[y][x] == 'L' && occupied == 0 {
				nseats[y] = nseats[y][:x] + "#" + nseats[y][x+1:]
			}
			if seats[y][x] == '#' && occupied >= 5 {
				nseats[y] = nseats[y][:x] + "L" + nseats[y][x+1:]
			}

		}
	}

	return nseats
}

func runPrint(lines []string) {
	for _, line := range lines {
		log.Printf("%v", line)
	}
}

//CountAndRunSeats runs the seats
func CountAndRunSeats(lines []string, alg int) int {
	equal := false

	oldSeats := doCopy(lines)

	count := 0
	for !equal {
		newSeats := runAlgorithm(oldSeats)
		if alg == 1 {
			newSeats = runAlgorithm2(oldSeats)
			if count > 1000 {
				runPrint(oldSeats)
				runPrint(newSeats)
				log.Fatalf("TRYING")
			}
		}

		tEq := true
		for i := range oldSeats {
			if newSeats[i] != oldSeats[i] {
				tEq = false
			}
		}

		equal = tEq
		oldSeats = doCopy(newSeats)

		count++
	}

	count = 0
	for _, row := range oldSeats {
		for _, char := range row {
			if char == '#' {
				count++
			}
		}
	}
	return count
}
