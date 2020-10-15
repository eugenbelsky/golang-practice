package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputINT uint = 1234
	bitCount := CountBits(inputINT)
	fmt.Printf("%v", bitCount)

}

func CountBits(v uint) int {
	a := strconv.FormatUint(uint64(v), 2)
	return strings.Count(a, strconv.Itoa(1))
}
