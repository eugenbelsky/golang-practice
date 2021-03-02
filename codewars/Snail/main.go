package main

import (
	"fmt"
)

func main() {

	playground := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	output := Snail(playground)
	fmt.Println(output)
}
func Snail(snaipMap [][]int) []int {

	var output []int

	output = append(output, snaipMap[0]...) // to the right
	snaipMap = snaipMap[1:]

	for i := 0; i < len(snaipMap[0]); i++ { // right
		output = append(output, snaipMap[0][i])
		if len(snaipMap[0]) > 1 {
			snaipMap[0] = snaipMap[0][1:]
		} else {
			snaipMap = snaipMap[1:]
		}
		i--
	}

	for i := 0; i < len(snaipMap); i++ { //down
		output = append(output, snaipMap[i][len(snaipMap[i])-1])
		if len(snaipMap[i]) > 1 {
			snaipMap[i] = snaipMap[i][:len(snaipMap[i])-1]
		} else {
			if len(snaipMap) > 1 {
				snaipMap = append(snaipMap[:i], snaipMap[:i+1]...)
			} else {
				return output
			}
		}

	}

	for i := len(snaipMap[len(snaipMap)-1]) - 1; i >= 0; i-- { // left
		output = append(output, snaipMap[len(snaipMap)-1][i])
		snaipMap[len(snaipMap)-1] = snaipMap[len(snaipMap)-1][:i-1]
	}

	for i := len(snaipMap) - 1; i >= 0; i-- { //up
		output = append(output, snaipMap[i][0])
		snaipMap[i] = snaipMap[i][0:]
	}

	fmt.Println(snaipMap)

	// for i,v :=range snaipMap{
	// 	ii,vv := range snaipMap[]{

	// 	}
	// }
	return output
}

func removeElement(slice []int, i int) []int{
	if len(slice) > 0 {
		if i == len(slice)-1 { //if last
			slice = slice[:len(slice)-1]
		} else {
			slice = append(slice[:i], slice[i+1:]..)
		}

	} 
	
	return slice
}
