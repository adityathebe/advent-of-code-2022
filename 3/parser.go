package main

import (
	"bufio"
	"strconv"
	"strings"
)

func cleanChar(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "[", "")
	s = strings.ReplaceAll(s, "]", "")
	return s
}

func parseInput(s *bufio.Scanner) Instruction {
	var (
		inst            Instruction
		charsList       [][]string
		isReadingStacks = true
	)

	for s.Scan() {
		line := s.Text()

		if isReadingStacks && line == "" {
			isReadingStacks = false
			continue
		}

		if isReadingStacks {
			var chars []string
			var buff string
			for _, char := range line {
				if len(buff) == 4 {
					chars = append(chars, cleanChar(buff))
					buff = ""
				}

				buff = buff + string(char)
			}

			if len(buff) == 3 {
				chars = append(chars, cleanChar(buff))
				buff = ""
			}

			charsList = append(charsList, chars)
		} else {
			line = strings.ReplaceAll(line, "move ", "")
			line = strings.ReplaceAll(line, "from", "")
			line = strings.ReplaceAll(line, "to", "")

			var move Move
			numbers := strings.Split(line, " ")
			for i, n := range numbers {
				n = strings.TrimSpace(n)
				if i == 0 {
					move.Count, _ = strconv.Atoi(n)
				} else if i == 2 {
					move.Src, _ = strconv.Atoi(n)
				} else if i == 4 {
					move.Dest, _ = strconv.Atoi(n)
				}
			}

			inst.Moves = append(inst.Moves, move)
		}
	}

	// Convert chars list to stack
	charsList = charsList[:len(charsList)-1]
	inst.Stacks = make([][]string, len(charsList[0]))
	for _, chars := range charsList {
		var ctr int
		for _, char := range chars {
			if strings.TrimSpace(char) != "" {
				inst.Stacks[ctr] = append([]string{char}, inst.Stacks[ctr]...)
			}

			ctr++
		}
	}

	return inst
}
