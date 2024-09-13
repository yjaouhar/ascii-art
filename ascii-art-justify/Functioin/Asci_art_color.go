package Func

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Color(args []string) {

	width := 0
	secend_flage := ""
	option := args[0]    // flag
	inputtext := args[1] // text to coloring
	substr := ""         // sub string to be color in this text
	banner := "standard"

	if len(args) <= 5 && strings.HasPrefix(args[0], "--color="){

		if len(args) == 3 {
			if strings.HasPrefix(args[1], "--align=") {
				secend_flage =strings.TrimPrefix(args[1], "--align=")
				inputtext = args[2]
			} else {
				if !Checkbanner(args[2]) {
					substr = ""
					inputtext = args[1]
					banner = args[2]
				} else {
					substr = args[1]
					inputtext = args[2]
				}

			}

		} else if len(args) == 4 {
			if strings.HasPrefix(args[1], "--align=") {
				secend_flage =strings.TrimPrefix(args[1], "--align=")
				if !Checkbanner(args[3]) {
					inputtext = args[2]
					banner = args[3]
				} else {
					inputtext = args[3]
					substr = args[2]
				}
			} else {
				inputtext = args[2]
				substr = args[1]
				banner = args[3]
			}

		} else {
			if strings.HasPrefix(args[1], "--align=") {
				secend_flage =strings.TrimPrefix(args[1], "--align=")
				substr = args[2]
				inputtext = args[3]
				banner = args[4]
			}
		}

		if Checkbanner(banner) {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> something")
			return
		}
		if strings.HasPrefix(args[1], "--align=") {
			cmd := exec.Command("stty", "size")
			cmd.Stdin = os.Stdin
			si, _ := cmd.Output()
			size := strings.TrimSpace(string(si))
			parts := strings.Split(size, " ")
			width, _ = strconv.Atoi(parts[1])
		}

		Asci_art(option, inputtext, banner, substr, secend_flage, width)

	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> something")
		return
	}
}
