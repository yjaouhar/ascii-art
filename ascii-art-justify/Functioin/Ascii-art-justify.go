package Func

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	option      string
	secend_flage string
	inputtext   string
	substr      string
	banner      string
)

func Justify(arg []string) {
	if (len(arg) < 2 || len(arg) > 5) || !strings.HasPrefix(arg[0], "--align=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}
	option = strings.TrimPrefix(arg[0], "--align=")
	inputtext = arg[1]
	banner = "standard"
	if len(arg) == 2 && strings.HasPrefix(arg[1], "--color=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}
	if len(arg) == 3 {
		if strings.HasPrefix(arg[1], "--color=") {
			secend_flage = arg[1]
			inputtext = arg[2]
		} else {
			inputtext = arg[1]
			banner = arg[2]
			if Checkbanner(banner) {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
				return
			}
		}
	} else if len(arg) == 4 {
		if strings.HasPrefix(arg[1], "--color=") {
			secend_flage = arg[1]
			if !Checkbanner(arg[3]) {
				inputtext = arg[2]
				banner = arg[3]
			} else {
				substr = arg[2]
				inputtext = arg[3]
			}
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
			return
		}

	} else if len(arg) == 5 {
		if strings.HasPrefix(arg[1], "--color=") {
			secend_flage = arg[1]
			substr = arg[2]
			inputtext = arg[3]
			banner = arg[4]
			if Checkbanner(banner) {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
				return
			}
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")

			return
		}
	}

	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	si, _ := cmd.Output()
	size := strings.TrimSpace(string(si))
	parts := strings.Split(size, " ")
	width, _ := strconv.Atoi(parts[1])
	Asci_art(option, inputtext, banner, substr, secend_flage, width)
}
