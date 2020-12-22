package a2020

import (
	"fmt"
)

func buildMap(lines []string) map[int]map[int]map[int]bool {

	m := make(map[int]map[int]bool)

	for y := range lines {
		for x, ch := range lines[y] {
			m2, ok := m[y]
			if !ok {
				m2 = make(map[int]bool)
			}
			if ch == '#' {
				m2[x] = true
			}
			m[y] = m2
		}
	}

	mm := make(map[int]map[int]map[int]bool)
	mm[0] = m

	return mm
}

func buildMap2(lines []string) map[int]map[int]map[int]map[int]bool {

	m := make(map[int]map[int]bool)

	for y := range lines {
		for x, ch := range lines[y] {
			m2, ok := m[y]
			if !ok {
				m2 = make(map[int]bool)
			}
			if ch == '#' {
				m2[x] = true
			}
			m[y] = m2
		}
	}

	mm := make(map[int]map[int]map[int]bool)
	mm[0] = m
	mmm := make(map[int]map[int]map[int]map[int]bool)
	mmm[0] = mm

	return mmm
}

func computeACycle(x, y, z int, m map[int]map[int]map[int]bool) bool {
	neighbours := 0
	ncount := 0
	for nz := z - 1; nz <= z+1; nz++ {
		for ny := y - 1; ny <= y+1; ny++ {
			for nx := x - 1; nx <= x+1; nx++ {
				if nx != x || ny != y || nz != z {
					ncount++
					if m2, ok := m[nz]; ok {
						if m3, ok2 := m2[ny]; ok2 {
							if m3[nx] {
								neighbours++
							}
						}
					}
				}
			}
		}
	}

	if m[z][y][x] {
		return !(neighbours < 2 || neighbours > 3)
	}

	return neighbours == 3
}

func computeACycle2(x, y, z, w int, m map[int]map[int]map[int]map[int]bool) bool {
	neighbours := 0
	ncount := 0
	for nw := w - 1; nw <= w+1; nw++ {
		for nz := z - 1; nz <= z+1; nz++ {
			for ny := y - 1; ny <= y+1; ny++ {
				for nx := x - 1; nx <= x+1; nx++ {
					if nx != x || ny != y || nz != z || nw != w {
						ncount++
						if m3, ok := m[nw]; ok {
							if m2, ok := m3[nz]; ok {
								if m3, ok2 := m2[ny]; ok2 {
									if m3[nx] {
										neighbours++
									}
								}
							}
						}
					}
				}
			}
		}
	}

	if m[w][z][y][x] {
		return !(neighbours < 2 || neighbours > 3)
	}

	return neighbours == 3
}

func runCycle(m map[int]map[int]map[int]bool) map[int]map[int]map[int]bool {
	minx, miny, maxx, maxy, minz, maxz := 0, 0, 0, 0, 0, 0

	m2 := make(map[int]map[int]map[int]bool)
	for z := range m {
		for y := range m[z] {
			for x := range m[z][y] {
				if x < minx {
					minx = x
				}
				if x > maxx {
					maxx = x
				}
			}
			if y < miny {
				miny = y
			}
			if y > maxy {
				maxy = y
			}
		}
		if z < minz {
			minz = z
		}
		if z > maxz {
			maxz = z
		}
	}

	for nz := minz - 1; nz <= maxz+1; nz++ {
		m2[nz] = make(map[int]map[int]bool)
		for ny := miny - 1; ny <= maxy+1; ny++ {
			m2[nz][ny] = make(map[int]bool)
			for nx := minx - 1; nx <= maxx+1; nx++ {
				m2[nz][ny][nx] = computeACycle(nx, ny, nz, m)
			}
		}
	}

	return m2
}

func runCycle2(m map[int]map[int]map[int]map[int]bool) map[int]map[int]map[int]map[int]bool {
	minx, miny, maxx, maxy, minz, maxz, minw, maxw := 0, 0, 0, 0, 0, 0, 0, 0

	m2 := make(map[int]map[int]map[int]map[int]bool)
	for w := range m {
		for z := range m[w] {
			for y := range m[w][z] {
				for x := range m[w][z][y] {
					if x < minx {
						minx = x
					}
					if x > maxx {
						maxx = x
					}
				}
				if y < miny {
					miny = y
				}
				if y > maxy {
					maxy = y
				}
			}
			if z < minz {
				minz = z
			}
			if z > maxz {
				maxz = z
			}
		}
		if w < minw {
			minw = w
		}
		if w > maxw {
			maxw = w
		}
	}

	for nw := minw - 1; nw <= maxw+1; nw++ {
		m2[nw] = make(map[int]map[int]map[int]bool)
		for nz := minz - 1; nz <= maxz+1; nz++ {
			m2[nw][nz] = make(map[int]map[int]bool)
			for ny := miny - 1; ny <= maxy+1; ny++ {
				m2[nw][nz][ny] = make(map[int]bool)
				for nx := minx - 1; nx <= maxx+1; nx++ {
					m2[nw][nz][ny][nx] = computeACycle2(nx, ny, nz, nw, m)
				}
			}
		}
	}
	return m2
}

func printMap(m map[int]map[int]map[int]bool) {
	minx, miny, maxx, maxy, minz, maxz := 0, 0, 0, 0, 0, 0
	for z := range m {
		for y := range m[z] {
			for x := range m[z][y] {
				if x < minx {
					minx = x
				}
				if x > maxx {
					maxx = x
				}
			}
			if y < miny {
				miny = y
			}
			if y > maxy {
				maxy = y
			}
		}
		if z < minz {
			minz = z
		}
		if z > maxz {
			maxz = z
		}
	}
	for z := minz; z <= maxz; z++ {
		fmt.Printf("z=%v\n", z)
		for y := miny; y <= maxy; y++ {
			for x := minx; x <= maxx; x++ {
				if m[z][y][x] {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}

			}
			fmt.Printf("\n")
		}
	}
}

//RunLife runs the life iteration
func RunLife(lines []string, cycles int) int {
	m := buildMap(lines)
	//printMap(m)
	for i := 0; i < cycles; i++ {
		m2 := runCycle(m)
		m = m2
	}

	sumv := 0
	for x := range m {
		for y := range m[x] {
			for _, val := range m[x][y] {
				if val {
					sumv++
				}
			}
		}
	}

	return sumv
}

//RunLife2 runs the life iteration
func RunLife2(lines []string, cycles int) int {
	m := buildMap2(lines)
	for i := 0; i < cycles; i++ {
		m2 := runCycle2(m)
		m = m2
	}

	sumv := 0
	for x := range m {
		for y := range m[x] {
			for z := range m[x][y] {
				for _, val := range m[x][y][z] {
					if val {
						sumv++
					}
				}
			}
		}
	}

	return sumv
}
