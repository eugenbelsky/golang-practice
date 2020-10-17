package main

import (
	"fmt"
)

func main() {
	var userInput float32
	_, _ = fmt.Scan(&userInput)
	fmt.Printf("%v", int(userInput))
}
