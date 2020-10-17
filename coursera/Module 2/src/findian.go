package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	userPrompt := "Enter a string:\n"
	fmt.Printf("%v", userPrompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	inputString := scanner.Text()

	runeSlice := []rune(inputString)

	if !(strings.EqualFold(string(runeSlice[0]), "i")) {
		fmt.Println("Not Found!")
		return
	}

	if !(strings.EqualFold(string(runeSlice[len(runeSlice)-1]), "n")) {
		fmt.Println("Not Found!")
		return
	}

	var count int
	for _, v := range runeSlice {
		if strings.EqualFold(string(v), "a") {
			count = count + 1
			break
		}
	}
	if count == 0 {
		fmt.Println("Not Found!")
		return
	}

	fmt.Println("Found!")

}
