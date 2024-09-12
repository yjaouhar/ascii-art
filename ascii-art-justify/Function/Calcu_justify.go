package Func

import (
	"strings"
)

// calcule number of space to be printed
func Calc_justify(elm string, matrix [][]string, widthInt int, option string) string {
	spec := ""
	nbword := strings.Fields(elm)

	lentext := Count(nbword, elm, matrix, option)
	nbspec := (widthInt - lentext)
	l := len(nbword) - 1
	if option == "center" {
		nbspec = (widthInt - lentext) / 2
	} else if option == "justify" {
		if l <= 0 {
			l = 1
		}
		nbspec = ((widthInt - lentext) / l)
	}
	for i := 0; i < nbspec; i++ {
		spec += " "
	}
	return spec
}
