package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]

	var mirror string

	for _, word := range args {
		if word >= 'a' && word <= 'z' {
			mirror += string('z' - (word - 'a'))
		} else if word >= 'A' && word <= 'Z' {
			mirror += string('Z' - (word- 'A'))
		} else {
			mirror += string(mirror)
		}
	}
	for _, aplM := range mirror {
		z01.PrintRune(aplM)
	}
	z01.PrintRune('\n')
}
