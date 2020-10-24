package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	message := "Please enter up to 10 integers separated by whitespace:"
	InputLimit := 20

	UserInput := UserPrompt(&message)

	CheckInputLen(UserInput, InputLimit)

	samples := ParseInput(UserInput)

	BubbleSort(samples)

	for _, number := range samples {
		fmt.Printf("%d ", number)
	}

}

func BubbleSort(samples []int) {

	for SortWindow := len(samples) - 1; SortWindow > 0; SortWindow-- {
		SwapFlag := 0

		for index := 0; index < SortWindow; index++ {

			if samples[index] > samples[index+1] {
				Swap(samples, index)
				SwapFlag = 1
			}
		}

		if SwapFlag == 0 {
			return
		}
	}
}

func Swap(samples []int, index int) {

	tmp := samples[index]
	samples[index] = samples[index+1]
	samples[index+1] = tmp
	return
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

func CheckInputLen(UserInput []string, InputLimit int) {

	if len(UserInput) > InputLimit {
		log.Println("User input is too long!")
		os.Exit(1)
	}

	return

}
