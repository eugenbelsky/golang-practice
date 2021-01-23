package main

import (
	"fmt"
	"strings"
)

func main() {
	SpinWords("This is another test")

}

func SpinWords(str string) string {
	var words []string
	var letters []rune
	var lettersSpinned []rune
	var wordsSpinned []string

	words = strings.Fields(str)

	// for i := 0; i < len(words); i++ {
	// 	letters = []rune(words[i])
	// 	lettersSpinned = append(lettersSpinned, letters)
	// }
	for _, ch := range words {
		lettersSpinned = nil
		letters = []rune(ch)
		if len(letters) >= 5 {
			for i := (len(letters) - 1); i >= 0; i-- {
				lettersSpinned = append(lettersSpinned, letters[i])
			}
		} else {
			lettersSpinned = append(lettersSpinned, letters...)
		}

		wordsSpinned = append(wordsSpinned, string(lettersSpinned))
	}
	strSpinned := strings.Join(wordsSpinned, " ")
	fmt.Printf(strSpinned)
	return strSpinned
} // SpinWords

// notes
//   string to rune[]
//   for each rune[] :new rune[] append form last to first
//   new rune[] to string[]
//   string.Join.
//(i := len(letters); i >= ; i--)
