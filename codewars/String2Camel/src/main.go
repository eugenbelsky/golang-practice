package main

import (
	"strings"
	"unicode"
)

func main() {

	ToCamelCase("the-Stealth-warrior")
	ToCamelCase("The_Stealth_Warrior")
}

func ToCamelCase(s string) string {

	var words []string
	var letters []rune
	var camel []string

	if strings.Contains(s, "-") {

		words = strings.Split(s, "-")
		// fmt.Printf("%v", wordSLice)

	}

	if strings.Contains(s, "_") {

		words = strings.Split(s, "_")
		// fmt.Printf(wordSlice[0])

	}

	for i, ch := range words {

		letters = []rune(ch)

		if (unicode.IsUpper(letters[0]) == false) && i >= 1 {

			letters[0] = unicode.ToUpper(letters[0])

		}

		camel = append(camel, string(letters))

	}

	camelForm := strings.Join(camel, "")

	// fmt.Printf("%v", camel)

	// Convert to runes
	// for wordSlice[0]; if runeSlice[0].isLower  => skip; else  runeSlice[0] => Upper
	// for i = 1, i++, wordslice [i] {
	//		runeSlice[0] -> to upper
	//		newS[].append
	//	}
	//
	return camelForm
}
