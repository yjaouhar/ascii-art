package Func

import (
	"fmt"
	"os"
	"os/exec"
)

// create a command take url of banner and download
func Curl(banner string) bool {
	url := "https://learn.zone01oujda.ma/git/root/public/raw/branch/master/subjects/ascii-art/" + banner // url of benner file
	cmd := exec.Command("curl", "-o", banner, url) // creates a new Cmd struct to execute the curl command.
	_, er := cmd.Output()
	if er != nil {
		fmt.Println("cant open this file")
		os.Exit(1)
	}
	return true
}
