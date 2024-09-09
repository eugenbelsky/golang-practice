package main

import (
	"flag"
	"fmt"
	"strings"
)

type arrayFlags []int

func (i *arrayFlags) String() string {
	str := make([]string)
	return strings.Join(i, " ")
}
func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var fieldFlags arrayFlags

func main() {
	flag.Var(&fieldFlags, "f", "field to cut")
	filepath := flag.Arg(0)
	flag.Parse()

	fmt.Println(*fieldArg)
	fmt.Println(filepath)

}
