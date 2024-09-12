package Func

import (
	"strings"
)

// ascii-art normal
func Asci_art(option string, inputtext string, banner string) {
	if IsEmpty(inputtext) {
		if !Is_print(inputtext) {
			return
		}
		inputsep := strings.Split(inputtext, "\\n")
		if strings.ReplaceAll(inputtext, "\\n", "") == "" && len(inputtext) > 1 {
			inputsep = inputsep[1:]
		}
		contentfile := Readfile(banner)

		matrix := Split((string(contentfile)), banner)
		Print(inputsep, matrix, option)

	}
}
