package main

import "testing"

func TestPartA(t *testing.T) {
	records := []struct {
		Input  string
		Output int
	}{
		{Input: "bvwbjplbgvbhsrlpgdmjqwftvncz", Output: 5},
		{Input: "nppdvjthqldpwncqszvftbrmjlhg", Output: 6},
		{Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", Output: 10},
		{Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", Output: 11},
	}

	for i, td := range records {
		output := findMarker(td.Input, 4)
		if output != td.Output {
			t.Fatalf("[%d]. Expected %d Got %d", i, td.Output, output)
		}
	}
}

func TestPartB(t *testing.T) {
	records := []struct {
		Input  string
		Output int
	}{
		{Input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", Output: 19},
		{Input: "bvwbjplbgvbhsrlpgdmjqwftvncz", Output: 23},
		{Input: "nppdvjthqldpwncqszvftbrmjlhg", Output: 23},
		{Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", Output: 29},
		{Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", Output: 26},
	}

	for i, td := range records {
		output := findMarker(td.Input, 14)
		if output != td.Output {
			t.Fatalf("[%d]. Expected %d Got %d", i, td.Output, output)
		}
	}
}
