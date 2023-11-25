package main

import (
	"fmt"
	"os"
	"strings"
)

func countVowels(inputString string) int {
	vowels := "aeiou"
	counter := 0

	for _, str := range inputString {
		if strings.Contains(vowels, string(str)) {
			counter += 1
		}
	}

	return counter
}

func containsDuplicates(inputString string) bool {
	for index, str := range inputString {
		if index + 1 < len(inputString) && string(str) == string(inputString[index + 1]) {
			return true
		}
	}

	return false
}

func containsRestricted(inputString string) bool {
	return strings.Contains(inputString, "ab") ||
		strings.Contains(inputString, "cd") ||
		strings.Contains(inputString, "pq") ||
		strings.Contains(inputString, "xy") 
}

func containsPairs(inputString string) bool {
	for i := 0; i < len(inputString) - 1; i++ {
		pair := string(inputString[i:(i+2)])

		if strings.Count(inputString, pair) > 1 {
			return true
		}
	}

	return false
}

func containsUwu(inputString string) bool {
	for index, str := range inputString {
		if index + 2 < len(inputString) && string(str) == string(inputString[index + 2]) {
			return true
		}
	}

	return false
}

func main() {
	dat, _ := os.ReadFile("input.txt")
	g := strings.Split(string(dat), "\n")
	var counter = 0

	for _, string := range g {
		// part 1
		// if countVowels(string) > 2 && containsDuplicates(string) && !containsRestricted(string) {
		// part 2
		if containsPairs(string) && containsUwu(string) {
			counter += 1
		}
	}

	fmt.Println(counter)
}