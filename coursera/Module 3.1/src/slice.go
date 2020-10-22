package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	s := make([]int, 3) //init empty n3 slice

	for {

		userPrompt := "Enter an int:\n" //  Print to user an input request
		fmt.Printf("%v", userPrompt)

		scanner := bufio.NewScanner(os.Stdin) // Collect an input with the scanner
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		t := scanner.Text()
		if t == "X" {
			quic
			break
		}

		v, err := strconv.Atoi(scanner.Text()) // Convert input to int
		if err != nil {
			fmt.Println("Not a proper int input: %v", err)
			continue
		}

		s = append(s, v) // Append the value to the slice

		sort.Ints(s) // Sort the slice and print
		fmt.Println(s)

	}
}
