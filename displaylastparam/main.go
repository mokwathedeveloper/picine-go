package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <=1 {
		return
	}
		lastarg := os.Args[1:]
		for _, arg1 := range lastarg[len(lastarg)-1] {
			z01.PrintRune(arg1)
		}
		z01.PrintRune('\n')
	}
