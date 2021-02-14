package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Word struct {
	name    string
	credits int
}

func main() {
	topCredit := solve("zodiacs")
	fmt.Println(topCredit)
}

func solve(str string) int {
	var re = regexp.MustCompile("(a|e|i|o|u)")
	parsedInput := re.ReplaceAllString(str, " ")
	topWord := High(parsedInput)
	return topWord
}
func High(s string) int {
	var topWord Word
	wordsChart := CreateChart(s)
	for _, v := range wordsChart {
		if v.credits > topWord.credits {
			topWord = v
		}
	}
	return topWord.credits
}
func CreateChart(s string) []Word {
	var wordsChart []Word
	words := strings.Fields(s)
	for _, v := range words {
		var creditSum int
		for _, z := range []rune(v) {
			ltCredit := int(z) - 96 // utf-8 -> byte starts from a=97 ...
			creditSum = creditSum + ltCredit

		}
		word := Word{name: v, credits: creditSum}
		wordsChart = append(wordsChart, word)
	}
	return wordsChart
}
