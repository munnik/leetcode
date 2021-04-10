package main

import "fmt"

func compress(chars []byte) int {
	currentPosition := 0
	previousPosition := 0
	var previousChar byte

	for currentPosition < len(chars) {
		if currentPosition == 0 {
			previousChar = chars[currentPosition]
			currentPosition++
			continue
		}

		if previousChar == chars[currentPosition] {
			currentPosition++
			continue
		}

		if currentPosition-previousPosition == 1 {
			// only keep one character. no modification to chars required
			previousChar = chars[currentPosition]
			previousPosition = currentPosition
			currentPosition++
		} else {
			// keep one character and a number
			numbers := []byte(fmt.Sprintf("%d", currentPosition-previousPosition))
			chars = append(
				append(
					chars[0:previousPosition+1],
					numbers...,
				),
				chars[currentPosition:]...,
			)
			currentPosition = previousPosition + 1 + len(numbers)
			previousChar = chars[currentPosition]
			previousPosition = currentPosition
		}
	}

	if currentPosition-previousPosition > 1 {
		// keep one character and a number
		numbers := []byte(fmt.Sprintf("%d", currentPosition-previousPosition))
		chars = append(
			chars[0:previousPosition+1],
			numbers...,
		)
	}

	return len(chars)
}
