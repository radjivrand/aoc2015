package main

import (
	"fmt"
	"crypto/md5"
	"strconv"
	"bytes"
)

func main() {
	startingString := "bgvyzdsv"
	// reference := []byte{0,0,16} // part 1
	reference := []byte{0,0,0,1} // part 2

	for i := 1; i < 10000000; i++ {
		b := []byte(startingString + strconv.Itoa(i))

		res := md5.Sum(b)
		test := res[0:3]

		if bytes.Compare(test, reference) == -1  {
			fmt.Println(i)
			break
		}
	}
}