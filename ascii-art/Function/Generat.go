package Func

import "fmt"

// print the txte
func Generat(Text []string, matrix [][]string) {
	for j := 0; j < len(Text); j++ {
		//	check if an element in input text is empty add newline
		if Text[j] == "" {
			fmt.Println()
			continue

		}
		for i := 0; i < 8; i++ { // print 8 line
			for k := 0; k < len(Text[j]); k++ {
				fmt.Print(matrix[int(Text[j][k]-32)][i]) // print the asci-art
			}
			fmt.Println()
		}

	}
}
