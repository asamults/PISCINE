package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for ch := 'a'; ch <= 'z'; ch++ {
		z01.PrintRune(ch)
		z01.PrintRune('\n')
	}
	// Corrected: Use fmt.Println to print a newline
}
