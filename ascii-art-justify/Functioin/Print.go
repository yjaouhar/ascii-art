package Func

import (
	"fmt"
	"os"
	"strings"
)

// print the sscii-art
func Print(elm string, matrix [][]string, option string, color string, substr string, color_flage string, nbspec int) string {
	print := ""
	writ := ""
	spac := ""
	for i := 0; i < nbspec; i++ {
		spac += " "
	}
		fmt.Println(spac)
	for i := 0; i < 8; i++ { // The lines that will be written the ascii-art
		if option == "center" || option == "right" {
			print += spac
		}
		for k := 0; k < len(elm); k++ {
			if (k != 0 && k != len(elm)-1) && (elm[k] == ' ' && option == "justify") {
				print += (spac)
			}
			for k != len(elm)-1 && (elm[k] == ' ' && option == "justify") {
				k++
			}
			if strings.HasPrefix(option, "--color=") || color_flage != "" {
				if substr != "" {
					if strings.HasPrefix(elm[k:], substr) { // serche for the sub string in the text
						for x := 0; x < len(substr); x++ {
							print += (color + matrix[int(elm[k+x]-32)][i] + "\033[0m") // print the asci-art
						}
						k += len(substr) - 1

					} else { // if the sub string don't trouv in text
						print += ("\033[0m" + matrix[int(elm[k]-32)][i])
					}
				} else {
					print += (color + matrix[int(elm[k]-32)][i] + "\033[0m")
				}
			} else {
				print += (matrix[int(elm[k]-32)][i])
			}
		}
		print += "\n"
	}

	if strings.HasPrefix(os.Args[1], "--output") && len(os.Args) > 2 {
		writ += print
	} else {
		fmt.Print(print) // print the asci-art
	}
	return writ
}
