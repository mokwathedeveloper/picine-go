package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// numAgrs := len(os.Args[1]) - 1
	// numstr := []rune{}

	// if numAgrs <= 0 {
	// 	numstr = append(numstr, '0')
	// }
	// for numAgrs > 1 {
	// 	numstr = append([]rune{rune(numAgrs%10 + '0')}, numstr...)
	// 	numAgrs /= 10
	// }
	// for _, count := range numstr {
	// 	z01.PrintRune(count)
	// }
	// z01.PrintRune('\n')

if len(os.Args) <=1 {
	z01.PrintRune('0')
	z01.PrintRune('\n')
	return
}

args := os.Args[1:]
count := '0'

for range args {
	count++
}
z01.PrintRune(count)
z01.PrintRune('\n')

}
