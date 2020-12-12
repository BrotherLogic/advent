package a2020

import (
	"strconv"
)

//MoveWaypoint make the moves
func MoveWaypoint(moves []string) int {
	wx := 10
	wy := 1
	posx := 0
	posy := 0
	for _, move := range moves {
		vl32, _ := strconv.ParseInt(move[1:], 10, 32)
		val := int(vl32)
		switch move[0] {
		case 'N':
			wy += val
		case 'S':
			wy -= val
		case 'E':
			wx += val
		case 'W':
			wx -= val
		case 'L':
			for val > 270 {
				val -= 360
			}
			switch val {
			case 90:
				wx, wy = -wy, wx
			case 180:
				wx, wy = -wx, -wy
			case 270:
				wx, wy = wy, -wx
			}
		case 'R':
			for val > 270 {
				val -= 360
			}
			switch val {
			case 90:
				wx, wy = wy, -wx
			case 180:
				wx, wy = -wx, -wy
			case 270:
				wx, wy = -wy, wx
			}
		case 'F':
			posx += val * (wx)
			posy += val * (wy)
		}
	}

	return abs(posx) + abs(posy)
}

//MakeMoves make the moves
func MakeMoves(moves []string) int {
	dir := 90

	posx := 0
	posy := 0
	for _, move := range moves {
		vl32, _ := strconv.ParseInt(move[1:], 10, 32)
		val := int(vl32)
		switch move[0] {
		case 'N':
			posy += val
		case 'S':
			posy -= val
		case 'E':
			posx += val
		case 'W':
			posx -= val
		case 'L':
			dir -= val
		case 'R':
			dir += val
		case 'F':
			switch dir {
			case 0:
				posy += val
			case 90:
				posx += val
			case 180:
				posy -= val
			case 270:
				posx -= val
			}
		}

		for dir > 270 {
			dir -= 360
		}
		for dir < 0 {
			dir += 360
		}
	}

	return abs(posx) + abs(posy)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
