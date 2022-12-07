package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	from int
	to   int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	ans := findOverlappingPairs(scanner)
	fmt.Println(ans)
}

func findOverlappingPairs(scanner *bufio.Scanner) int {
	var totalOverlaps int
	for scanner.Scan() {
		line := scanner.Text()
		assignments := parseLine(line)
		if isPairOverlapping(assignments[0], assignments[1]) {
			totalOverlaps++
		} else {
			fmt.Println(assignments)
		}
	}

	return totalOverlaps
}

func sortPair(a, b Assignment) (Assignment, Assignment) {
	pairWithSmallerStart := a
	pairWithLargerStart := b
	if b.from < a.from || (b.from == a.from && b.to > a.to) {
		pairWithSmallerStart = b
		pairWithLargerStart = a
	}

	return pairWithSmallerStart, pairWithLargerStart
}

func isPairOverlapping(a, b Assignment) bool {
	smallerPair, largerPair := sortPair(a, b)
	return smallerPair.to >= largerPair.to
}

func parseLine(line string) []Assignment {
	pairs := strings.Split(line, ",")

	assigments := make([]Assignment, 0, len(pairs))
	for _, pair := range pairs {
		ranges := strings.Split(pair, "-")

		var a Assignment
		a.from, _ = strconv.Atoi(ranges[0])
		a.to, _ = strconv.Atoi(ranges[1])
		assigments = append(assigments, a)
	}

	return assigments
}
