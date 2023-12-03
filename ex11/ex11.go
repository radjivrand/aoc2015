package main

import (
	"fmt"
	s "strings"
)

var input = "vzbxxyzz"

func main() {
	for {
		input = getNext(input)

		if !validateForbidden(input) {
			continue
		}

		if !validateDoubles(input) {
			continue
		}

		if !validateConsecutive(input) {
			continue
		}

		break
	}

	fmt.Println(input)
}

func addOne(runeslice []rune, index int) []rune {
	runeslice[index] += 1
	if runeslice[index] == 123 {
		runeslice[index] = 97
		runeslice = addOne(runeslice, index - 1)
	}

	return runeslice
}

func getNext(input string) string {
	r := []rune(input)
	length := len(r)

	return string(addOne(r, length - 1))
}

func validateForbidden(input string) bool {
	return !(s.Contains(input, "i") || s.Contains(input, "o") || s.Contains(input, "l"))
}

func validateDoubles(input string) bool {
	r := []rune(input)

	var first rune
	for i := 0; i < len(r) - 1; i++ {
		if r[i] == r[i+1] {
			if first == 0 {
				first = r[i+1]
				continue
			} else if r[i] != first {
				return true
			}
		}
	}

	return false
}

func validateConsecutive(input string) bool {
	r := []rune(input)

	for i := 0; i < len(r) - 3; i++ {
		if r[i] + 1 == r[i+1] && r[i+1] + 1 == r[i+2] {
			return true
		}
	}

	return false
}