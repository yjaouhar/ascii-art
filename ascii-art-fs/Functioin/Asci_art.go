package Func

import (
	"strings"
)

// ascii-art normal
func Asci_art(args []string) {
	inputtext := args[0]
	if IsEmpty(inputtext) {
		Is_print(inputtext)
		inputsep := strings.Split(inputtext, "\\n")
		if strings.ReplaceAll(inputtext, "\\n", "") == "" && len(inputtext) > 1 {
			inputsep = inputsep[1:]
		}
		contentfile := Readfile("standard")
		matrix := Split((string(contentfile)), "")
		Print(inputsep, matrix)

	}
}
