package Func

import "strings"

// splite standar file withe newline in a matrix
func Split(contentfile string) [][]string {
	matrix := [][]string{}
	result := strings.Split(contentfile[1:len(contentfile)-1], "\n\n") // any art character as an element in slice
	for x := 0; x < len(result); x++ {
		matrix = append(matrix, strings.Split(result[x], "\n")) //split element in slice in another slice
	}
	return matrix
}
