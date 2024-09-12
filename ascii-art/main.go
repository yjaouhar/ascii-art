package main

import (
	"os"
	"strings"

	Func "ascii-art/Function"
)

func main() {
	Func.CheckArgs(os.Args)
	inputtext := os.Args[1]
	if Func.IsEmpty(inputtext) {
		Func.Is_print(inputtext)
		inputsep := strings.Split(inputtext, "\\n")
		if strings.ReplaceAll(inputtext, "\\n", "") == "" && len(inputtext) > 1 {
			inputsep = inputsep[1:]
		}
		contentfile := Func.Readfile("standard.txt")

		matrix := Func.Split((string(contentfile)))
		Func.Generat(inputsep, matrix)

	}
}
