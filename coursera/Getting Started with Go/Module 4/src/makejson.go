package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	msgName := "Enter your name:"
	msgAddress := "Enter your address:"
	pMap := make(map[string]string)

	pMap["name"] = shPrompt(&msgName)
	pMap["address"] = shPrompt(&msgAddress)
	pJSON, _ := json.Marshal(pMap)

	fmt.Println(string(pJSON))

}

func shPrompt(msg *string) string {

	fmt.Println(*msg)

	scanner := bufio.NewScanner(os.Stdin) // Collect an input with the scanner
	scanner.Scan()
	return scanner.Text()

}
