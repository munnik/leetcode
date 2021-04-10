package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte("a"), []byte("a")},
		{[]byte("aabbccc"), []byte("a2b2c3")},
		{[]byte("abbbbbbbbbbbb"), []byte("ab12")},
		{[]byte("aaabbaa"), []byte("a3b2a2")},
		{[]byte("aaaaaaaaaaaaaabbb"), []byte("a14b3")},
	}

	for i, testCase := range tests {
		t.Run(
			fmt.Sprintf("Test %d", i),
			func(t *testing.T) {
				got := compress(testCase.input)
				assert.Equal(t, got, len(testCase.want))
				assert.Equal(t, testCase.want, testCase.input[:got])
			},
		)
	}
}
