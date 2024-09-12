package Func

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Generate the ascii-art-justify
func Justify(args []string) {
	option := strings.TrimPrefix(args[0], "--align=")
	inputtext := args[1]
	banner := "standard"
	if len(args) > 2 {
		banner = args[2]
	}
	if len(args) > 3 || !(option == "center" || option == "left" || option == "right" || option == "justify") || (Checkbanner(banner)) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	} else {
		cmd := exec.Command("stty", "size")
		cmd.Stdin = os.Stdin
		si, _ := cmd.Output()
		size := strings.TrimSpace(string(si))
		parts := strings.Split(size, " ")
		width, _ := strconv.Atoi(parts[1])
		Asci_art(option, inputtext, banner, width)
	}
}
