package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('a')
		z01.PrintRune('\n')
		return
	}

	arg := os.Args[1]
	if len(arg) == 0 {
		z01.PrintRune('a')
		z01.PrintRune('\n')
		return
	}

	z01.PrintRune(rune(arg[0])) // Печатаем первый символ строки
	z01.PrintRune('\n')
}
