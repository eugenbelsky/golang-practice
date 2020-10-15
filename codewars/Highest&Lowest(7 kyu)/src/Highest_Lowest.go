package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputSTR string = "-1"
	fmt.Printf(HighAndLow(inputSTR))
}

func HighAndLow(in string) string {
	strSlice := strings.Split(in, " ")

	var numSlice = make([]int, len(strSlice))

	for i, v := range strSlice {
		numSlice[i], _ = strconv.Atoi(v)
	}

	var numMin int = numSlice[0]
	var numMax int = numSlice[0]
	for _, v := range numSlice {

		if v > numMax {
			numMax = v
		}
		if v < numMin {
			numMin = v
		}
	}
	return strconv.Itoa(numMax) + " " + strconv.Itoa(numMin)
}
