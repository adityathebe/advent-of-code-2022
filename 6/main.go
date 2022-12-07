package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(findMarker(string(data)))
}

func findMarker(datastream string) int {
	var buff Buff
	for i, char := range datastream {
		buff.Put(char)
		if !buff.IsFull() {
			continue
		}

		if buff.IsDistinct() {
			return i + 1
		}
	}

	return -1
}

type Buff struct {
	data []rune
}

func (t *Buff) Put(v rune) {
	t.data = append(t.data, v)
	if len(t.data) > 4 {
		t.data = t.data[1:]
	}
}

func (t *Buff) IsFull() bool {
	return len(t.data) == 4
}

func (t *Buff) IsDistinct() bool {
	var m = make(map[rune]struct{}, len(t.data))
	for _, char := range t.data {
		if _, ok := m[char]; ok {
			return false
		}

		m[char] = struct{}{}
	}

	return true
}
