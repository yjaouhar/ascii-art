package Func

import (
	"fmt"
	"strings"
)

func Color(args []string) {
	if len(args) < 2 || len(args) > 4 || !strings.HasPrefix(args[0], "--color=") { // checke len and flag
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> something")
		return
	}
	option := args[0]    // flag
	inputtext := args[1] // text to coloring
	substr := ""         // sub string to be color in this text
	banner := "standard"
	if len(args) > 2 {
		if !Checkbanner(args[2]) {
			substr = ""
			inputtext = args[1]
			banner = args[2]
		} else {
			substr = args[1]
			inputtext = args[2]
		}

		if len(args) > 3 {
			banner = args[3]
			if (!Checkbanner(args[2]) && len(args) > 3  ) || Checkbanner(banner){
				fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> something")
				return
			}
			
		}
	}

	Asci_art(option, inputtext, banner, substr, "", 0)
}
