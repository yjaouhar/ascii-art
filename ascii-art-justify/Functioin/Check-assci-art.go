package Func

import (
	"fmt"
	"strings"
)

// check ascii-art if ascii-art or ascii-art-fs or ascii-art-output or ascii-art-color
func Check_ascii_art(args []string) {
	if len(args) < 1 { // don't have option or string
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> something")
		return

	} else {
		inputtext := args[0]
		banner := "standard"
		if len(args) == 1 && !strings.HasPrefix(args[0], "--output") && !strings.HasPrefix(args[0], "--color") && !strings.HasPrefix(args[0], "--align") {
			Asci_art("", inputtext, banner, "", "", 0)
			return
		} else {
			if strings.HasPrefix(args[0], "--output") { // verife ascii-art-output
				Output(args)
			} else if strings.HasPrefix(args[0], "--color") { // verife ascii-art-color
				Color(args)
			} else if strings.HasPrefix(args[0], "--align") {
				Justify(args)
			} else {
				banner := args[1]
				if len(args) > 2 || Checkbanner(banner) { // verife ascii-art-fs
					fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
					return
				}
				Asci_art("", inputtext, banner, "", "", 0)
			}
		}
	}
}
