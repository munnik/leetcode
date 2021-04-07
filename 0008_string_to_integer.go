package main

func myAtoi(s string) int {
	numberMap := map[rune]int32{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

	var result int32
	var resultBefore int32

	negative := false
	spaceAllowed := true
	signAllowed := true

	for _, c := range s {
		number, isNumberFound := numberMap[c]
		switch {
		case c == ' ' && spaceAllowed:
			continue
		case c == '-' && signAllowed:
			signAllowed = false
			spaceAllowed = false

			negative = true
		case c == '+' && signAllowed:
			signAllowed = false
			spaceAllowed = false
		case isNumberFound:
			signAllowed = false
			spaceAllowed = false

			resultBefore = result
			result *= 10
			if negative {
				result -= number
				if resultBefore < result || resultBefore != result/10 {
					return -2147483648
				}
			} else {
				result += number
				if resultBefore > result || resultBefore != result/10 {
					return 2147483647
				}
			}
		default:
			return int(result)
		}
	}

	return int(result)
}
