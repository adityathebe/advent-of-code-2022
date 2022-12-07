package main

import (
	"strconv"
	"testing"
)

func Test_isPairOverlapping(t *testing.T) {
	tests := []struct {
		a    Assignment
		b    Assignment
		want bool
	}{
		{a: Assignment{4, 4}, b: Assignment{4, 4}, want: true},
		{a: Assignment{1, 5}, b: Assignment{1, 4}, want: true},
		{a: Assignment{1, 4}, b: Assignment{1, 5}, want: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := isPairOverlapping(tt.a, tt.b); got != tt.want {
				t.Errorf("isPairOverlapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
