package main

import (
	"fmt"
	"strings"
)

var morseMap = map[string]byte{
	".-":     'A',
	"-...":   'B',
	"-.-.":   'C',
	"-..":    'D',
	".":      'E',
	"..-.":   'F',
	"--.":    'G',
	"....":   'H',
	"..":     'I',
	".---":   'J',
	"-.-":    'K',
	".-..":   'L',
	"--":     'M',
	"-.":     'N',
	"---":    'O',
	".--.":   'P',
	"--.-":   'Q',
	".-.":    'R',
	"...":    'S',
	"-":      'T',
	"..-":    'U',
	"...-":   'V',
	".--":    'W',
	"-..-":   'X',
	"-.--":   'Y',
	"--..":   'Z',
	".----":  '1',
	"..---":  '2',
	"...--":  '3',
	"....-":  '4',
	".....":  '5',
	"-....":  '6',
	"--...":  '7',
	"---..":  '8',
	"----.":  '9',
	"-----":  '0',
	".-.-.-": '.',
	"--..--": ',',
	"..--..": '?',
	"-.-.--": '!',
	"-....-": '-',
	"-..-.":  '/',
	".--.-.": '@',
	"-.--.":  '(',
	"-.--.-": ')',
}

var morseMapException = map[string]string{
	"...---...": "SOS",
	".---.":     "ROGER",
	"---.---":   "OVER",
	"........":  "ERROR",
}

func main() {

	encoded := "... --- ..."

	decoded := DecodeMorse(encoded)
	fmt.Println(decoded)

}

func DecodeBits(bits string) string {
	return ""
}

func DecodeMorse(sentence string) string {

	sentence = strings.TrimSpace(sentence)  // remove trailing spaces
	words := strings.Split(sentence, "   ") // separate words

	var output []string

	for _, v := range words { // decode O(n)^2
		runes := strings.Fields(v)
		var word []string
		for _, r := range runes {
			if _, ok := morseMap[r]; ok { // if basic character
				word = append(word, string(morseMap[r]))
				continue
			}

			if _, ok := morseMapException[r]; ok { // if basic character
				word = append(word, morseMapException[r])
				continue
			}

			word = append(word, morseMapException["........"]) // if not found -> ERROR

		}
		output = append(output, strings.Join(word, ""))
	}

	return strings.Join(output, " ")

}
