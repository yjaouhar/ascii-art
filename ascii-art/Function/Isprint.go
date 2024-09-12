package Func

import (
	"fmt"
	"os"
)

// check the  argument is have a caracter printible
func Is_print(s string) bool {
	slr := []rune(s)
	for i := 0; i < len(slr); i++ {
		if !(slr[i] >= 32 && slr[i] <= 126) {
			fmt.Println("ERROR : The is not printibale caracter")
			os.Exit(1)
			return false
		}
	}
	return true
}
