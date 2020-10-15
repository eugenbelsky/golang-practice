package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputINT uint = 1234
	bitCount := CountBits(inputINT)
	fmt.Printf("%v", bitCount)
}

func CountBits(v uint) int {
	a, _ := strconv.Atoi(strconv.FormatUint(uint64(v), 2))
	for i := range a {

	}

	return a
}
