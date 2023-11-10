package mock

import (
	"strings"
)

func Palindrome(input string) bool {
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")
	sliceChar := strings.Split(input, "")

	for i := 0; i < len(sliceChar)/2; i++ {
		if sliceChar[i] != sliceChar[len(sliceChar)-1-i] {
			return false
		}
	}

	return true
}
