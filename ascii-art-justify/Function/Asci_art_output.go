package Func

import (
	"fmt"
)

// ascii-art-output
func Output(args []string) {
	option := args[0]
	inputtext := args[1]
	banner := "standard"
	if len(args) > 2 {
		banner = args[2]
	}
	if len(args) > 3 || Checkbanner(banner) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	} else {
		Asci_art(option, inputtext, banner, 0)
	}
}
