package Func

import (
	"fmt"
	"os"
	"strings"
)

// writ the ascii-art in file
func Writ_in_file(print string, option string) {
	filename := strings.TrimPrefix(option, "--output=")
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error : can't craet the is file")
		return
	}
	_, err = file.WriteString(print)
	if err != nil {
		fmt.Println("can't writ in the is file")
		return
	}
}
