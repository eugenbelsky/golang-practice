package main

import (
	"fmt"
	"sort"
)

func main() {
	digits := []int{2, 1, 5, 3}
	mostFrequent := HighestRank(digits)
	fmt.Printf("%v", mostFrequent)
}

func HighestRank(nums []int) int {

	var mostFrequent int
	sort.Ints(nums)
	counter := 0
	maxCounter := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			counter++
			if counter >= maxCounter {
				maxCounter = counter
				mostFrequent = nums[i+1]
			}
		} else {
			counter = 0
		}
	}

	if maxCounter == 0 {
		return nums[len(nums)-1]
	} else {
		return mostFrequent
	}
}
