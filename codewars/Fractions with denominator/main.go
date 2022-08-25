package main

import "fmt"

func main() {
	d := 10000123
	fractions := ProperFractions(d)
	fmt.Println(fractions)
}

// func ProperFractions(d int) int { //SLOW On^2
// 	seq := GenerateSeq(d)
// 	counter := 0
// OUTER:
// 	for i := range seq {

// 		for n := range seq[0 : i+1] {
// 			if seq[i]%seq[n] == 0 && d%seq[n] == 0 && seq[n] != 1 {

// 				continue OUTER

// 			}
// 		}
// 		counter++
// 	}
// 	return counter
// }

func ProperFractions(d int) int {
	divs := make(map[int]bool)

	seq := GenerateSeq(d)

	for _, v := range seq { // get simple dividers
		if d%v == 0 && v != 1 {
			divs[v] = true
		}
	}

	for key := range divs { // get complex dividers
		b := key
		i := 2
		for b*i < d {
			divs[b*i] = true
			i++
		}
	}

	counter := 0
	for _, v := range seq {
		_, ok := divs[v]
		if !ok {

			counter++
		}

	}
	return counter
}

func GenerateSeq(d int) []int {
	seq := make([]int, d-1) // fill out range of nominators
	for i := range seq {
		seq[i] = i + 1
	}
	return seq

}
