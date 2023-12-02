package main

import (
	"fmt"
	"os"
	"strings"
	"slices"
	"strconv"
)

var ref [][]string
var stars, permutations []string
var numbers string

func main() {
	dat, _ := os.ReadFile("input.txt")
	g := strings.Split(string(dat), "\n")

	for _, row := range g {
		splitted := strings.Split(row, " = ")
		cities := strings.Split(splitted[0], " to ")

		if !slices.Contains(stars, cities[0]) {
			stars = append(stars, cities[0])
		}

		if !slices.Contains(stars, cities[1]) {
			stars = append(stars, cities[1])
		}

		splatted := strings.Split(row, " = ")

		s := []string{cities[0], cities[1], splatted[1]}
		ref = append(ref, s)
	}

	for i := 0; i < len(stars); i++ {
		numbers += strconv.Itoa(i)
	}

	permute(numbers, 0, len(numbers) - 1)

	var smallest = 10000000
	var biggest = 0
	var smallestPerm = ""
	for _, perm := range permutations {
		var distance = 0

		r := []rune(perm)
		for i := 0; i < len(r) - 1; i++ {
			if i < len(r) - 1 {
				firstIndex, _ := strconv.Atoi(string(r[i]))
				nextIndex, _ := strconv.Atoi(string(r[i + 1]))
				distance += findDistance(stars[firstIndex], stars[nextIndex])
			}
		}

		//part1
		if distance < smallest {
			smallest = distance
			smallestPerm = perm
		}

		//part2
		if distance > biggest {
			biggest = distance
		}
	}

	// fmt.Println(smallest)
	fmt.Println(smallestPerm)
	fmt.Println(biggest)
}

func findDistance(first, second string) int {
	for _, instruction := range ref {
		if (instruction[0] == first && instruction[1] == second) ||
		(instruction[1] == first && instruction[0] == second) {
			res,_ := strconv.Atoi(instruction[2])
			return res
		}
	}
	return 0
}

func permute(inputString string, startIndex, endIndex int) {
	if startIndex == endIndex {
		permutations = append(permutations, inputString)
	} else {
		for i := startIndex; i <= endIndex; i++ {
			r := []rune(inputString)
			r[startIndex], r[i] = r[i], r[startIndex]
			inputString = string(r)

			permute(inputString, startIndex + 1, endIndex)

			u := []rune(inputString)
			u[startIndex], u[i] = u[i], u[startIndex]
			inputString = string(u)
		}
	}
}