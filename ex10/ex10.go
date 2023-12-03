package main

import (
	"fmt"
	"strconv"
)

func lookAndSay(input string) string {
	r := []rune(input)
	var results []rune
	var next rune

	counter := 0
	for index, char := range r {
		if len(r) > index + 1 || index == len(r) - 1{
			if len(r) > index + 1 {
				next = r[index + 1]
			}

			counter += 1

			if char != next || index == len(r) - 1 {
				counterString := strconv.Itoa(counter)
				counterSlice := []rune(counterString)

				results = append(results, counterSlice...)
				results = append(results, char)

				counter = 0
			}
		}
	}

	return string(results)
}

func main() {
	// 11131221232112111312212221121123222113 too high –– :D length was asked!

	// var input = "1"
	var input = "3113322113"

	for i := 0; i < 50; i++ {
		input = lookAndSay(input)
	}

	fmt.Println(len(input))
}