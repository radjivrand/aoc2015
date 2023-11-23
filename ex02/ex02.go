package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func getSmallestSideArea(a, b, c int) (int, int) {
	var ab = a * b
	var bc = b * c
	var ac = a * c

	if ab <= bc && ab <= ac {
		return a, b
	} else if bc <= ab && bc <= ac {
		return b, c
	} else {
		return a, c
	}
}

func main() {
	var sum = 0
	var ribbon = 0
	dat, _ := os.ReadFile("input.txt")

	g := strings.Split(string(dat), "\n")

	for _, dimensions := range g {
		var test = strings.Split(dimensions, "x")
		var x, _ = strconv.Atoi(test[0])
		var y, _ = strconv.Atoi(test[1])
		var z, _ = strconv.Atoi(test[2])
		var smallA, smallB = getSmallestSideArea(x, y, z)

		sum += 2 * x * y  + 2 * y * z + 2* x * z + smallA * smallB
		ribbon += x * y * z + 2 * (smallA + smallB)
	}

	fmt.Println(sum)
	fmt.Println(ribbon)
}