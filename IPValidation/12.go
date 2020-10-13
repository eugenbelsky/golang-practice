package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "010"
	if _, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%q looks like a number.\n", v)
	}

}
