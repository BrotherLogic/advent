package main

import (
	"math/rand"
	"testing"
	"time"
)

var day4Part1Table = []struct {
	in   []string
	out  int
	out2 int
}{
	{[]string{
		"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:05] falls asleep",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 00:55] wakes up",
		"[1518-11-01 23:58] Guard #99 begins shift",
		"[1518-11-02 00:40] falls asleep",
		"[1518-11-02 00:50] wakes up",
		"[1518-11-03 00:05] Guard #10 begins shift",
		"[1518-11-03 00:24] falls asleep",
		"[1518-11-03 00:29] wakes up",
		"[1518-11-04 00:02] Guard #99 begins shift",
		"[1518-11-04 00:36] falls asleep",
		"[1518-11-04 00:46] wakes up",
		"[1518-11-05 00:03] Guard #99 begins shift",
		"[1518-11-05 00:45] falls asleep",
		"[1518-11-05 00:55] wakes up",
	}, 240, 4455}}

func TestDay4Part1(t *testing.T) {
	for _, tt := range day4Part1Table {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		narr := []string{}
		for _, i := range r.Perm(len(tt.in)) {
			narr = append(narr, tt.in[i])
		}

		v1, v2 := computeSleep(narr)
		if v1 != tt.out || v2 != tt.out2 {
			t.Errorf("Error in sleep computation: %v,%v but was %v, %v", tt.out, tt.out2, v1, v2)
		}
	}
}
