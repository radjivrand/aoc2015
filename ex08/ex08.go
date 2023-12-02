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
	var res = 0;


	for _, string := range g {
		// part 1
		// converted,_ := strconv.Unquote(string);
		// res += len(string);
		// res -= len(converted);

		// part 2
		converted := strconv.Quote(string);
		res += len(converted);
		res -= len(string);
	}

	fmt.Println(res);
}
