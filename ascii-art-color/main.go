package main

import (
	"os"

	Func "ascii-art/Functioin"
)

func main() {
	Func.Check_ascii_art(os.Args[1:])
}
