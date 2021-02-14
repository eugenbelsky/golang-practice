package main

import "fmt"

func main() {
	parens := "hi(hi))("
	isValid := ValidParentheses(parens)
	fmt.Printf("%v", isValid)

}
func ValidParentheses(parens string) bool {
	counter := 0
	characters := []rune(parens)
	for _, v := range characters {
		if string(v) == "(" {
			counter++
		}
		if string(v) == ")" {
			counter--
			if counter < 0 {
				return false
			}

		}
	}

	if counter == 0 {
		return true
	} else {
		return false
	}

}
