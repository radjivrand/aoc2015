package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type rule struct {
	a, op, b, value, dest string
	applied bool
}

func findValue(inputString string, results map[string]int) int {
	var intA int
	intA, err := strconv.Atoi(inputString)
	if err != nil {
		var present bool
		intA, present = results[inputString]
		if !present {
			intA = 7777777
		} else {
			intA = results[inputString]
		}
	}

	return intA
}

func main() {
	dat, _ := os.ReadFile("input.txt")
	g := strings.Split(string(dat), "\n")
	rules := make([]rule, len(g))
	finalMap := make(map[string]int)

	for index, string := range g {
		rules[index] = parseRule(string)
	}

	for j := 0; j < 1000; j++ {
		for i := 0; i < len(rules); i++ {
			ptrRule := &rules[i]

			binA := findValue(ptrRule.a, finalMap)
			binB := findValue(ptrRule.b, finalMap)
			binValue := findValue(ptrRule.value, finalMap)

			if !ptrRule.applied {
				switch ptrRule.op {
				case "AND":
					if binA == 7777777 || binB == 7777777 {
						continue
					}
					finalMap[ptrRule.dest] = binA & binB
				case "OR":
					if binA == 7777777 || binB == 7777777 {
						continue
					}
					finalMap[ptrRule.dest] = binA | binB
				case "LSHIFT":
					if binA == 7777777 || binB == 7777777 {
						continue
					}
					finalMap[ptrRule.dest] = binA << binB
				case "RSHIFT":
					if binA == 7777777 || binB == 7777777 {
						continue
					}
					finalMap[ptrRule.dest] = binA >> binB
				case "NOT":
					if binB == 7777777 {
						continue
					}
					finalMap[ptrRule.dest] = ^binB
				case "":
					if binValue == 7777777 {
						continue
					}
					finalMap[ptrRule.dest] = binValue
				}

				ptrRule.applied = true
			}
		}
	}

	fmt.Println(finalMap)
}

func parseRule(line string) rule {
	splitted := strings.Split(line, " -> ")
	var op, a, b, value string

	firstPart := strings.Split(splitted[0], " ")
	switch len(firstPart) {
	case 1:
		value = firstPart[0]
	case 2:
		op = firstPart[0]
		b = firstPart[1]
	case 3:
		a = firstPart[0]
		op = firstPart[1]
		b = firstPart[2]
	}

	rule := rule{a, op, b, value, splitted[1], false}

	return rule
}