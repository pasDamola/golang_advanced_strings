package main

import "fmt"

func main() {
	fmt.Printf("You can set the width of the printed number, left aligned: %5d|\n", 10)
	// You can set the width of the printed number, left
	// aligned:    10|
	fmt.Printf("You can make numbers right-aligned with a given width: %-5d|\n", 10)
}
