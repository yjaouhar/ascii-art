package Func

import "fmt"

// print the sscii-art
func Print(inputsep []string, matrix [][]string, option string) {
	print := ""
	for j := 0; j < len(inputsep); j++ {
		if inputsep[j] == "" { // newline leaves empty string so when he find empty string print newline
			print += "\n"
			continue
		}
		for i := 0; i < 8; i++ { // The lines that will be written the ascii-art
			for k := 0; k < len(inputsep[j]); k++ {
				print += (matrix[int(inputsep[j][k]-32)][i])
			}
			print += "\n"
		}
	}
	if option != "" {
		Writ_in_file(print, option) // writ the  ascii-art in file
	} else {
		fmt.Print(print) // print the asci-art
	}
}
