package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
if len(os.Args) < 1 {
	return
}
firstArgs := os.Args[1:]

for _, first := range firstArgs[0]{
	z01.PrintRune(first)
}
	z01.PrintRune('\n')
}