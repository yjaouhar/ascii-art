package Func

import (
	"fmt"
	"strings"
)

// check ascii-art if ascii-art or ascii-art-fs or ascii-art-output
func Check_ascii_art(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: go run . [STRING] \n\nEX: go run . something ")
		return

	} else {
		inputtext := args[0]
		banner := "standard"
		if len(args) == 1 {
			Asci_art("", inputtext, banner, 0)
			return
		}
		if len(args) > 1 {
			if strings.HasPrefix(args[0], "--output") { // if args 1 have a flage (--output)
				Output(args)
			} else if strings.HasPrefix(args[0], "--align") {
				Justify(args)
			} else {
				banner := args[1]
				if len(args) > 2 || !(banner == "standard" || banner == "shadow" || banner == "thinkertoy") {
					fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
					return
				}
				Asci_art("", inputtext, banner, 0)
			}
		}
	}
}
