package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {

	a := PromtValue("Enter a:")
	v := PromtValue("Enter v:")
	s := PromtValue("Enter s:")

	fn := GenDisplaceFn(a, v, s)

	t := PromtValue("Enter t:")

	fmt.Printf("S = %g \n", fn(t))
}

func PromtValue(stdout string) float64 {

	var stdin string

	fmt.Println(stdout)

	fmt.Scan(&stdin)

	ConvValue, err := strconv.ParseFloat(stdin, 64)

	if err != nil {
		log.Fatal(err)
	}

	return ConvValue
}

func GenDisplaceFn(a, v, s float64) func(t float64) float64 {

	fn := func(time float64) float64 {

		return 0.5*a*math.Pow(time, 2) + v*time + s
	}

	return fn

}
