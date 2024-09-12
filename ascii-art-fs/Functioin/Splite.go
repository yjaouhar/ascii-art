package Func

import (
	"strings"
)

// splite the file withe newline in a matrix
func Split(contenctfile string, banner string) [][]string {
	matrix := [][]string{}
	if banner == "thinkertoy" {
		contenctfile = strings.Replace(contenctfile, "\r\n", "\n", -1)
	}
	result := strings.Split(contenctfile[1:len(contenctfile)-1], "\n\n") // any art character as an element in slice
	for x := 0; x < len(result); x++ {
		matrix = append(matrix, strings.Split(result[x], string('\n'))) // split element in slice in another slice
	}
	return matrix
}
