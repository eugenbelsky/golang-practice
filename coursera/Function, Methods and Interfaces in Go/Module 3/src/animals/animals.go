package animals

import (
	"fmt"
)

type Animal struct {
	Food       string
	Locomotion string
	Sound      string
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
