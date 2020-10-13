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
