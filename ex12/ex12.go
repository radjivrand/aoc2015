package main

import (
	"fmt"
	"os"
	// "strings"
	// "slices"
	// "strconv"
	"encoding/json"
)


func main() {
	dat, _ := os.ReadFile("input_test.txt")
	// g := strings.Split(string(dat), "\n")
	var s []interface{}

	g := json.Unmarshal(dat, &s)

	fmt.Println(g)
}

