package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		input []string
		want  [][]string
	}{
		{
			[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
			[][]string{
				{"root/a/1.txt", "root/c/3.txt"},
				{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"},
			},
		},
		{
			[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"},
			[][]string{
				{"root/a/1.txt", "root/c/3.txt"},
				{"root/a/2.txt", "root/c/d/4.txt"},
			},
		},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := findDuplicate(testCase.input)
				assert.Equal(t, testCase.want, got)
			},
		)
	}
}
