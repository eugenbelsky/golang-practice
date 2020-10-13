package main

import (
	"fmt"
	// "strconv"
	"strings"
)

func main() {
	var inputStr = "192.168.1.1"
	var a = strings.SplitAfter(inputStr, ".")
	// var IPArray = [4]uint8{"192"}
	for s := range a {
		fmt.Printf(a[s])
	}
	// grades := [3]{13, 145, 54}
	// fmt.Printf("Grades: %v", grades)

}

//12312123
// func is_valid_ip(ip string) bool {
//     var IPArray [4]uint8
//     a :=
// 	return false
// }
