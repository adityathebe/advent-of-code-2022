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

	fmt.Println(findMarker(string(data), 4))
	fmt.Println(findMarker(string(data), 14))
}

func findMarker(datastream string, buffSize int) int {
	buff := Buff{
		size: buffSize,
	}

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
	size int
}

func (t *Buff) Put(v rune) {
	t.data = append(t.data, v)
	if len(t.data) > t.size {
		t.data = t.data[1:] // Discard the first character
	}
}

func (t *Buff) IsFull() bool {
	return len(t.data) == t.size
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
