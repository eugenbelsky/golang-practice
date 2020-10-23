package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {

	var srtSlice []name
	var filename string

	fmt.Println("Enter file name:")

	fmt.Scan(&filename)

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		v := strings.Fields(scanner.Text())
		for _, k := range v {

			if len(k) > 20 {
				log.Fatal(k + " - value is too big(20+)")
				return
			}

		}
		srtSlice = append(srtSlice, name{fname: v[0], lname: v[1]})
	}

	for _, v := range srtSlice {
		fmt.Println("First Name:" + v.fname + " " + "Last Name:" + v.lname)
	}

}
