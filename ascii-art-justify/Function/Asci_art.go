package Func

import (
	"fmt"
	"os"
	"strings"
)

// Generate the  ascii-art
func Asci_art(option string, inputtext string, banner string, width int) {
	writ := ""
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
		spec := ""
		for _, elm := range inputsep {
			if strings.HasPrefix(os.Args[1], "--align") {
				if elm == "" {
					fmt.Println()
				}else{
					spec = Calc_justify(elm, matrix, width, option)
				Print(elm, matrix, option, spec)

				}
			} else {
				if elm == "" {
					writ += "\n"
				}else{
					writ += Print(elm, matrix, option, spec)
				}
				
			}
		}
		if writ != "" {
			Writ_in_file(writ, option) // writ the  ascii-art in file
		}
	}
}
