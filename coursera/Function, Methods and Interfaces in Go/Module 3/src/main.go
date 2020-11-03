package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Animal struct {
	Food       string
	Locomotion string
	Sound      string
}

var cow = Animal{
	Food:       "grass",
	Locomotion: "walk",
	Sound:      "moo",
}

var bird = Animal{
	Food:       "worms",
	Locomotion: "fly",
	Sound:      "peep",
}

var snake = Animal{
	Food:       "mice",
	Locomotion: "slither",
	Sound:      "hsss",
}

func main() {

	for {
		animal, method := PromtValue(">")

		m := strings.Title(method) // caused by interface call issue in MethodByName

		switch animal {

		case "snake":

			reflect.ValueOf(&snake).MethodByName(m).Call(nil)

		case "bird":

			reflect.ValueOf(&bird).MethodByName(m).Call(nil)

		case "cow":
			reflect.ValueOf(&cow).MethodByName(m).Call(nil)

		}
	}

}

func PromtValue(stdout string) (string, string) {

	var animal string
	var method string

	fmt.Printf(stdout)

	fmt.Scan(&animal)

	fmt.Scan(&method)

	return animal, method
}

func (a *Animal) Eat() {
	fmt.Printf("Food: %s \n", a.Food)
}

func (a *Animal) Move() {
	fmt.Printf("Locomotion method: %s \n", a.Locomotion)
}

func (a *Animal) Speak() {
	fmt.Printf("Sound: %s \n", a.Sound)
}
