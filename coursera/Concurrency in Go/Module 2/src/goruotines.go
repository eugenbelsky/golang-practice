package main

import (
	"fmt"
)

var c int = 10

func main() {

	c := 10
	go calc()
	go calc()
	fmt.Println(c)

}

func calc() {
	c = c + 2
}
