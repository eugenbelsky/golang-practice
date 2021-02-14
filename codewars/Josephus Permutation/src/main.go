package main

import (
	"fmt"
)

func main() {
	items := []interface{}{}
	// items := []interface{}{1, 2, 3, 4, 5, 6, 7}
	dead := Josephus(items, 3)
	fmt.Printf("%v", dead)
}

func Josephus(alive []interface{}, k int) []interface{} {

	var dead []interface{}
	counter := 0
	for {
		if len(alive) == 0 {
			break
		}
		for i := 0; i <= len(alive)-1; i++ {
			counter = counter + 1
			if counter == k {
				if i == len(alive)-1 {
					dead = append(dead, alive[i])
					alive = alive[:i]
				} else {
					dead = append(dead, alive[i])
					alive = append(alive[:i], alive[i+1:]...)
					i--
				}
				counter = 0
			}
		}
	}
	return dead
}
