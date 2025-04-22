package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	for ch := 'a'; ch <= 'z'; ch++ {
		z01.PrintRune(ch)
	}
	fmt.Println() // Corrected: Use fmt.Println to print a newline
}
