package main

import (
	"fmt"
)

func main() {
	digits := []int{2, 1, 5, 3, 1, 5}
	oddOccur := FindOdd(digits)
	fmt.Printf("%v", oddOccur)

}

func FindOdd(seq []int) int {
	digitMap := make(map[int][]int)
	var oddOccur int
	for i := 0; i < len(seq); i++ {
		digitMap[seq[i]] = append(digitMap[seq[i]], seq[i])
	}

	for key, value := range digitMap {
		if len(value)%2 != 0 {
			oddOccur = key
		}
	}

	return oddOccur
}
