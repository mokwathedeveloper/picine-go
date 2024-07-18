package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	i := 0
	for _, matched := range str2 {
		if  i < len(str1) && matched == rune(str1[i]) {
			i++
		}
	}
	if i == len(str1) {
		for _, matched := range str2 {
			z01.PrintRune(matched)
		}
	}
	z01.PrintRune('\n')

}
