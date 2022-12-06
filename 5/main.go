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

func (t *Instruction) ExecuteNew() string {
	for _, move := range t.Moves {
		var itr int
		var carrying []string
		for itr < move.Count {
			itr++
			var val string
			val, t.Stacks[move.Src-1] = pop(t.Stacks[move.Src-1])
			carrying = append([]string{val}, carrying...)
		}

		t.Stacks[move.Dest-1] = append(t.Stacks[move.Dest-1], carrying...)
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

	// crateMover9000 := inst.Execute()
	// fmt.Println(crateMover9000)

	crateMover9001 := inst.ExecuteNew()
	fmt.Println(crateMover9001)
}
