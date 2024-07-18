package main

import (
	"github.com/01-edu/z01"

	//"piscine"
)

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}

func LastRune(s string) rune {
	runes := []rune(s)
	index := len(runes) - 1
	return runes[index]
}