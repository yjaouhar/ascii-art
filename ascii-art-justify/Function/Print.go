package Func

import (
	"fmt"
	"os"
	"strings"
)

// print the sscii-art
func Print(elm string, matrix [][]string, option string, spec string) string {
	print := ""
	writ := ""
	elm = strings.TrimSpace(elm)
	for i := 0; i < 8; i++ { // The lines that will be written the ascii-art
		if option == "center" || option == "right" {
			print += (spec)
		}
		for k := 0; k < len(elm); k++ {
			if (k != 0 && k != len(elm)-1) && (elm[k] == ' ' && option == "justify") {
				print += (spec)
			}
			for k != len(elm)-1 && (elm[k] == ' ' && option == "justify") {
				k++
			}
			print += (matrix[int(elm[k]-32)][i])
		}
		print += "\n"
	}

	if !strings.HasPrefix(os.Args[1], "--output") {
		fmt.Print(print) // print the asci-art
		// fmt.Println()
	} else {
		writ += print
	}
	return writ
}
