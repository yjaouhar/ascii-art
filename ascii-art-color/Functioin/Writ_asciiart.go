package Func

import (
	"fmt"
	"os"
	"strings"
)

// writ the ascii-art in file
func Writ_asciiart(containoutput string, option string) {
	filename := strings.TrimPrefix(option, "--output=")

	if !(strings.HasSuffix(filename, ".txt")) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error : can't craet the is file")
		return
	}
	_, err = file.WriteString(containoutput)
	if err != nil {
		fmt.Println("can't writ in the is file")
		return
	}
}
