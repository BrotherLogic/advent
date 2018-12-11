package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	children []node
	metadata []int
}

func buildTree(arr []int, start int) (node, int) {
	nNode := node{make([]node, arr[start]), make([]int, arr[start+1])}
	pointer := start + 2
	for i := range nNode.children {
		tree, newPointer := buildTree(arr, pointer)
		nNode.children[i] = tree
		pointer = newPointer
	}

	for i := range nNode.metadata {
		nNode.metadata[i] = arr[pointer+i]
	}

	return nNode, pointer + len(nNode.metadata)
}

func sumMeta(nNode node) int {
	metaSum := 0
	for i := range nNode.children {
		metaSum += sumMeta(nNode.children[i])
	}

	for _, v := range nNode.metadata {
		metaSum += v
	}

	return metaSum
}

func sumRoot(nNode node) int {
	if len(nNode.children) == 0 {
		sumv := 0
		for _, v := range nNode.metadata {
			sumv += v
		}

		return sumv
	} else {
		sumv := 0
		for _, v := range nNode.metadata {
			if v-1 < len(nNode.children) {
				sumv += sumRoot(nNode.children[v-1])
			}
		}
		return sumv
	}
}

func computeSum(arr []int) (int, int) {
	treeRoot, _ := buildTree(arr, 0)
	return sumMeta(treeRoot), sumRoot(treeRoot)
}

func solveDay8(file string) {
	arr := []int{}
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		elems := strings.Split(scanner.Text(), " ")
		for _, v := range elems {
			iv, _ := strconv.Atoi(v)
			arr = append(arr, iv)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	v1, v2 := computeSum(arr)
	fmt.Printf("Day 8 [1]. %v\n", v1)
	fmt.Printf("Day 8 [2]. %v\n", v2)
}
