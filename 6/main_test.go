package main

import "testing"

func TestSome(t *testing.T) {
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
		output := findMarker(td.Input)
		if output != td.Output {
			t.Fatalf("[%d]. Expected %d Got %d", i, td.Output, output)
		}
	}
}
