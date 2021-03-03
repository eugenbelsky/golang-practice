package main

import (
	"fmt"
)

type Playground struct {
	snaipMap [][]int
}

func main() {

	initmap := [][]int{}
	output := Snail(initmap)
	fmt.Println(output)
}
func Snail(snaipMap [][]int) []int {

	output := []int{} // empty slice instead of nil

	var playground Playground
	playground.snaipMap = snaipMap

	for len(playground.snaipMap) > 0 {
		playground.Right(&output)
		playground.Down(&output)
		playground.Left(&output)
		playground.Up(&output)
	}

	return output
}

func (playground *Playground) Up(output *[]int) {
	if len(playground.snaipMap) == 0 {
		return
	}
	for i := len(playground.snaipMap) - 1; i >= 0; i-- { //up
		*output = append(*output, playground.snaipMap[i][0])

		if len(playground.snaipMap[i]) > 1 {
			playground.snaipMap[i] = removeElement(playground.snaipMap[i], 0)
		} else {
			playground.snaipMap = removeRaw(playground.snaipMap, i)
		}
	}
	return
}

func (playground *Playground) Down(output *[]int) {

	if len(playground.snaipMap) == 0 {
		return
	}
	for i := 0; i < len(playground.snaipMap); i++ { //down
		*output = append(*output, playground.snaipMap[i][len(playground.snaipMap[i])-1])

		if len(playground.snaipMap[i]) > 1 {
			playground.snaipMap[i] = removeElement(playground.snaipMap[i], len(playground.snaipMap[i])-1)
		} else {
			playground.snaipMap = removeRaw(playground.snaipMap, i)
			i--
		}

	}
	return
}

func (playground *Playground) Right(output *[]int) {

	if len(playground.snaipMap) == 0 {
		return
	}
	for i := 0; i < len(playground.snaipMap[0]); i++ { // right
		*output = append(*output, playground.snaipMap[0][i])
		playground.snaipMap[0] = removeElement(playground.snaipMap[0], i)
		i--
	}
	if len(playground.snaipMap[0]) == 0 {
		playground.snaipMap = removeRaw(playground.snaipMap, 0)
	}
	return
}

func (playground *Playground) Left(output *[]int) {

	if len(playground.snaipMap) == 0 {
		return
	}
	for i := len(playground.snaipMap[len(playground.snaipMap)-1]) - 1; i >= 0; i-- { // left
		*output = append(*output, playground.snaipMap[len(playground.snaipMap)-1][i])
		playground.snaipMap[len(playground.snaipMap)-1] = removeElement(playground.snaipMap[len(playground.snaipMap)-1], i)
	}

	if len(playground.snaipMap[len(playground.snaipMap)-1]) == 0 {
		playground.snaipMap = removeRaw(playground.snaipMap, len(playground.snaipMap)-1)
	}

	return
}

func removeElement(slice []int, i int) []int {
	if len(slice) > 0 {
		if i == len(slice)-1 { //if last
			slice = slice[:len(slice)-1]
		} else {
			slice = append(slice[:i], slice[i+1:]...)
		}

	}
	return slice
}

func removeRaw(slice [][]int, i int) [][]int {
	if len(slice) > 0 {
		if i == len(slice)-1 { //if last
			slice = slice[:len(slice)-1]
		} else {
			slice = append(slice[:i], slice[i+1:]...)
		}

	}
	return slice
}
