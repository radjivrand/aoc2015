package main

import (
	"fmt"
	"os"
)

func main() {
	floor := 0

	dat, _ := os.ReadFile("input.txt")

	for idx, rune := range dat {
		if string(rune) == "(" {
			floor += 1
		} else {
			floor -= 1
		}

		if floor < 0 {
			fmt.Println(idx + 1)
			return
		}
	}

	fmt.Println(floor)
}