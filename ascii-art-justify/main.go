package main

import (
	"os"

	Func "justify/Function"
)

func main() {
	Func.Check_ascii_art(os.Args[1:])
}
