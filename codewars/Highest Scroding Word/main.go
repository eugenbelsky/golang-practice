package main

import (
	"fmt"
	"strings"
)

type Word struct {
	name    string
	credits int
}

func main() {

	// buf := make([]byte, 3)
	// letter := 'z'
	// utf8.EncodeRune(buf, letter)
	word := High("fnbcqvsday ouieapxt qxhgxvx frh bexrwfclff mfoq gub swphmxc vdfarmzhlr lllw wjpglt qqdgq")
	fmt.Println(word)

}
func High(s string) string {
	var topWord Word
	wordsChart := CreateChart(s)
	for _, v := range wordsChart {
		if v.credits > topWord.credits {
			topWord = v
		}
	}
	return topWord.name
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
		fmt.Println(v)
		fmt.Println(creditSum)
		word := Word{name: v, credits: creditSum}
		wordsChart = append(wordsChart, word)
	}
	return wordsChart
}
