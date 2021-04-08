package main

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"00000-42a1234", 0},
		{"-6147483648", -2147483648},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := myAtoi(testCase.input)
				if got != testCase.want {
					t.Fatalf("Got %v; want %v; with input %v", got, testCase.want, testCase.input)
				} else {
					t.Logf("Success!")
				}
			},
		)
	}
}
