package Func

import (
	"fmt"
	"image/color"
	"os"
	"strings"
)

// ascii-art normal
func Asci_art(option string, inputtext string, banner string, substr string, color_flage string, width int) {
	writ := "" // for writ this ascii-art in file
	namecolor := ""
	spec := 0
	if IsEmpty(inputtext) {
		if !Is_print(inputtext) {
			return
		}
		inputsep := strings.Split(inputtext, "\\n") // splite text
		if strings.ReplaceAll(inputtext, "\\n", "") == "" && len(inputtext) > 1 {
			inputsep = inputsep[1:]
		}
		contentfile := Readfile(banner)
		matrix := Split((string(contentfile)), banner)
		for _, elm := range inputsep {
			if strings.HasPrefix(os.Args[1], "--output") && len(os.Args) > 2 {
				if elm == "" {
					writ += "\n"
				} else {
					writ += Print(elm, matrix, option, "", "", color_flage, 0)
				}
			} else {
				colore := ""
				if strings.HasPrefix(os.Args[1], "--color=") {
					namecolor = strings.TrimPrefix(option, "--color=")
					sl, color_valide := Checkcolors(namecolor)
					if color_valide && len(sl) != 0 {
						code := color.RGBA{R: sl[0], G: sl[1], B: sl[2]}
						colore = Convertrgbtocode(code)

					} else {
						fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> something")
						return
					}
				} else if strings.HasPrefix(os.Args[1], "--align=") {

					color_flage = strings.TrimPrefix(color_flage, "--color=")
					sl, color_valide := Checkcolors(namecolor)
					if color_valide && len(sl) != 0 {
						code := color.RGBA{R: sl[0], G: sl[1], B: sl[2]}
						colore = Convertrgbtocode(code)

					} else {
						fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> something")
						return
					}
					spec = Calc_justify(elm, matrix, width, option)
				}
				if elm == "" {
					fmt.Println()
				} else {
					Print(elm, matrix, option, colore, substr, color_flage, spec)
				}
			}
		}
		if writ != "" {
			Writ_asciiart(writ, option) // writ the  ascii-art in file
		}
	}
}
