package main

import (
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	gameBoard *board
	currentGo *board
}

type board struct {
	value int
	next  *board
	prev  *board
}

func printBoard(g *game) string {
	str := ""
	rootVal := g.gameBoard.value
	boardRoot := g.gameBoard

	for len(str) == 0 || boardRoot.value != rootVal {
		if boardRoot.value == g.currentGo.value {
			str += fmt.Sprintf("(%v) ", boardRoot.value)
		} else {
			str += fmt.Sprintf("%v ", boardRoot.value)
		}
		boardRoot = boardRoot.next
	}

	return str
}

func play(g *game, val int) int {
	score := 0

	if val%23 == 0 {
		score += val
		counter := g.currentGo.prev.prev.prev.prev.prev.prev.prev
		score += counter.value

		// Remove the stone from play
		currPrev := counter.prev
		currNext := counter.next
		currPrev.next = currNext
		currNext.prev = currPrev

		// Turn ends
		g.currentGo = currNext
		return score
	}

	addboard := g.currentGo.next
	currNext := addboard.next
	currPrev := addboard

	newBoard := &board{value: val, next: currNext, prev: currPrev}
	currNext.prev = newBoard
	currPrev.next = newBoard

	g.currentGo = newBoard

	return score
}

func runGame(def string, winner int) (int, int) {
	elems := strings.Fields(def)

	numplayers, _ := strconv.Atoi(elems[0])
	last, _ := strconv.Atoi(elems[6])

	players := make([]int, numplayers)

	initialBoard := &board{value: 0}
	initialBoard.next = initialBoard
	initialBoard.prev = initialBoard
	g := &game{initialBoard, initialBoard}

	for i := 1; i <= last; i++ {
		score := play(g, i)
		players[(i-1)%numplayers] += score
	}

	if winner >= 0 {
		return players[winner], winner
	}

	maxv := 0
	index := 0
	for i, score := range players {
		if score > maxv {
			maxv = score
			index = i
		}
	}
	return maxv, index
}

func solveDay9() {
	answer, _ := runGame("438 players; last marble is worth 71626 points", -1)
	fmt.Printf("Day 9 [1]. %v\n", answer)
	part2, _ := runGame("438 players; last marble is worth 7162600 points", -1)
	fmt.Printf("Day 9 [2]. %v\n", part2)
}
