package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// imitate input
	var IPCheck = Is_valid_ip("1.124.10")
	fmt.Printf("%v\n", IPCheck)
	// parse input to array by dot division
	// a := []string(strings.SplitN(inputStr, ".", -1))
	// fmt.Printf("%v\n", a)
	// var IPArray [4]uint64
	// for i, s := range a {
	// 	var err error
	// 	IPArray[i], err = strconv.ParseUint(s, 10, 8)
	// 	if err != nil {
	// 		fmt.Printf("false input")
	// 	}
	// 	fmt.Printf("%v\n", IPArray[i])
	// 	fmt.Printf("%v\n", s)
	// 	fmt.Printf("%v\n", err)
	// }
	// fmt.Printf("%v\n", IPArray)
	// grades := [3]{13, 145, 54}
	// fmt.Printf("Grades: %v", grades)

}

func Is_valid_ip(ip string) bool {
	a := []string(strings.SplitN(ip, ".", -1))
	var IPArray [4]uint64
	if len(a) != 4 {
		return false
	}
	for i, s := range a {
		var err error
		IPArray[i], err = strconv.ParseUint(s, 10, 8)
		if err != nil {
			return false
		}

	}

	return true
}
