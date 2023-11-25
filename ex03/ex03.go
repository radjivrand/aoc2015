package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	dat, _ := os.ReadFile("input.txt")
	g := strings.Split(string(dat), "")

	santasMap := make(map[string]int)
	var currentSantaXY [2]int
	var currentRoboXY [2]int

	currentSantaXY[0] = 0
	currentSantaXY[1] = 0
	currentRoboXY[0] = 0
	currentRoboXY[1] = 0

	santasMap[strconv.Itoa(currentSantaXY[0]) + "_" +  strconv.Itoa(currentSantaXY[1])] += 1

	for index, dir := range g {
		switch dir {
		case ">":
			if index%2 == 0 {
				currentSantaXY[0] += 1
			} else {
				currentRoboXY[0] += 1
			}
		case "<":
			if index%2 == 0 {
				currentSantaXY[0] -= 1
			} else {
				currentRoboXY[0] -= 1
			}
		case "^":
			if index%2 == 0 {
				currentSantaXY[1] -= 1
			} else {
				currentRoboXY[1] -= 1
			}
		case "v":
			if index%2 == 0 {
				currentSantaXY[1] += 1
			} else {
				currentRoboXY[1] += 1
			}
		}

		santasMap[strconv.Itoa(currentSantaXY[0]) + "_" +  strconv.Itoa(currentSantaXY[1])] += 1
		santasMap[strconv.Itoa(currentRoboXY[0]) + "_" +  strconv.Itoa(currentRoboXY[1])] += 1
	}

	fmt.Println(len(santasMap))
}