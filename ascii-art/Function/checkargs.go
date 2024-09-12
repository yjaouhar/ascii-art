package Func

import (
	"fmt"
	"os"
)

// check len of args
func CheckArgs(args []string) bool {
	if len(args) != 2 {
		fmt.Println("ERROR : You have the right tow argemunt ")
		os.Exit(1)
		return false
	}
	return true
}

// check input text is empty
func IsEmpty(inputtext string) bool {
	if inputtext != "" {
		return true
	}
	return false
}
