package Func

import "fmt"

// check ascii-art if ascii-art or ascii-art-fs
func Check_ascii_art(args []string) {
	if len(args) == 1 {
		Asci_art(args)
		return
	}
	if len(args) > 1 {
		banner := args[1]
		if (len(args) < 3) && (banner == "standard" || banner == "shadow" || banner == "thinkertoy") {
			Ascii_art_fs(args)
		} else {
			fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
			return
		}
	} else { // if len(args) == 0
		fmt.Println("Usage: go run . [STRING] \n\nEX: go run . something ")
		return
	}
}
