package main

import (
	"fmt"
	"math"
)

func main() {

	var slice1 []int
	// var slice2 []int

	// slice1 := nil
	slice2 := []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	valid := Comp(slice1, slice2)
	fmt.Println(valid)
}
func Comp(slice1 []int, slice2 []int) bool {

	if len(slice1) > 0 && len(slice2) > 0 {

	comp:
		for _, v := range slice1 {

			square := int(math.Pow(float64(v), 2))

			for _, z := range slice2 {
				if square == z {
					continue comp
				}
			}

			return false
		}

		return true
	}

	return false

}
