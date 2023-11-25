package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	dat, _ := os.ReadFile("input.txt")
	g := strings.Split(string(dat), "\n")

	var ledArr [1000][1000]bool

	for _, string := range g {
		interm := strings.Replace(string, "turn ", "turn", -1)
		interm = strings.Replace(interm, " through ", " ", -1)
		interm = strings.Replace(interm, ",", " ", -1)
		row := strings.Split(interm, " ")

		sX, _ := strconv.Atoi(row[1])
		sY, _ := strconv.Atoi(row[2])
		eX, _ := strconv.Atoi(row[3])
		eY, _ := strconv.Atoi(row[4])

		for i := sY; i <= eY; i++ {
			for j := sX; j <= eX; j++ {
				switch row[0] {
				case "turnon":
					ledArr[i][j] = true
				case "turnoff":
					ledArr[i][j] = false
				case "toggle":
					ledArr[i][j] = !ledArr[i][j]
				}
			}
		}
	}

	counter := 0

	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			if ledArr[i][j] {
				counter += 1
			}
		}
	}

	fmt.Println(counter)
}