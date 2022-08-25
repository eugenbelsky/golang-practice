package main

import (
	"fmt"
	"reflect"
)

func main() {
	word := "abba"
	words := []string{"aabb", "abcd", "bbaa", "dada"}
	anagrams := Anagrams(word, words)
	fmt.Println(anagrams)

}
func Anagrams(word string, words []string) []string {

	var anagrams []string
	checkWordMap := WordToMap(word)

	for _, w := range words {
		wordMap := WordToMap(w)
		if reflect.DeepEqual(checkWordMap, wordMap) {
			anagrams = append(anagrams, w)
		}
	}

	return anagrams

}

func WordToMap(word string) map[string]int {
	m := map[string]int{}
	for _, v := range word {
		if _, ok := m[string(v)]; ok {
			m[string(v)] = m[string(v)] + 1
		} else {
			m[string(v)] = 1
		}
	}
	return m
}
