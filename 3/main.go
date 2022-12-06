package main

import (
	"bufio"
	"fmt"
	"os"
)

type Move struct {
	Count int
	Src   int
	Dest  int
}

type Instruction struct {
	Stacks [][]string
	Moves  []Move
}

func pop(a []string) (string, []string) {
	x, a := a[len(a)-1], a[:len(a)-1]
	return x, a
}

func (t *Instruction) Execute() string {
	for _, move := range t.Moves {
		var itr int
		for itr < move.Count {
			itr++
			var val string
			val, t.Stacks[move.Src-1] = pop(t.Stacks[move.Src-1])
			t.Stacks[move.Dest-1] = append(t.Stacks[move.Dest-1], val)
		}
	}

	var result string
	for _, crate := range t.Stacks {
		if len(crate) == 0 {
			continue
		}

		last := crate[len(crate)-1]
		result += last
	}

	return result
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	inst := parseInput(scanner)
	answ := inst.Execute()
	fmt.Println(answ)
}
