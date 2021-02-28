package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	userOutput := "Enter integers divided by whitespace: "
	userInput := UserPrompt(&userOutput)
	samples := ParseInput(userInput)

	d1, d2 := divideSlice(samples)
	subSlice1, subSlice2 := divideSlice(d1)
	subSlice3, subSlice4 := divideSlice(d2)

	wg.Add(1)
	go sortByThread(&wg, &subSlice1)

	wg.Add(1)
	go sortByThread(&wg, &subSlice2)

	wg.Add(1)
	go sortByThread(&wg, &subSlice3)

	wg.Add(1)
	go sortByThread(&wg, &subSlice4)

	wg.Wait()

	result := []int{}
	result = append(append(subSlice1, subSlice2...), append(subSlice3, subSlice4...)...)
	sort.Ints(result)
	fmt.Printf("Final result: %v", result)

}

func divideSlice(samples []int) ([]int, []int) {
	var num = len(samples)

	mid := int(num / 2)

	var leftSubslice = make([]int, mid)
	var rightSubslice = make([]int, num-mid)

	for i := 0; i < num; i++ {
		if i < mid {
			leftSubslice[i] = samples[i]
		} else {
			rightSubslice[i-mid] = samples[i]
		}
	}

	return leftSubslice, rightSubslice
}

func sortByThread(wg *sync.WaitGroup, s *[]int) {

	fmt.Printf("This thread is sorting the following slice: %v \n", *s)
	sort.Ints(*s)
	wg.Done()

}

func UserPrompt(message *string) []string {

	fmt.Println(*message)

	scanner := bufio.NewScanner(os.Stdin) // Collect an input with the scanner

	scanner.Scan()

	UserInput := strings.Fields(scanner.Text())

	return UserInput
}

func ParseInput(UserInput []string) []int {

	var samples = []int{}

	for _, field := range UserInput {

		sample, err := strconv.Atoi(field)

		if err != nil {
			log.Fatal(err)
		}

		samples = append(samples, sample)

	}

	return samples

}
