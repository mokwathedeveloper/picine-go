package main

import "github.com/01-edu/z01"

func main() {
	// 	alpha := "aBcDeFgHiJkLmNoPqRsTuVwXyZ"

	// 	for _, sud := range alpha {
	// 		z01.PrintRune(sud)
	// 	}

	// z01.PrintRune('\n')
	for i := 'a'; i <= 'z'; i++ {
		if (i-'a')%2 == 0 {
			z01.PrintRune(i)
		} else {
			z01.PrintRune(i -32)
		}
	}
z01.PrintRune('\n')
}
