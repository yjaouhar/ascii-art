package Func

import (
	"strings"
)

// ascii-art-fs
func Ascii_art_fs(args []string) {
	banner := args[1]
	inputtext := args[0]

	if IsEmpty(inputtext) {
		Is_print(inputtext)
		inputsep := strings.Split(inputtext, "\\n")
		if strings.ReplaceAll(inputtext, "\\n", "") == "" && len(inputtext) > 1 {
			inputsep = inputsep[1:]
		}
		contentfile := Readfile(banner)
		matrix := Split((string(contentfile)), banner)
		Print(inputsep, matrix)

	}
}
