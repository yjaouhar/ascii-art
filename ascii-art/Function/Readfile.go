package Func

import (
	"fmt"
	"os"
	"os/exec"
)

// read a standar.txt file
func Readfile(filman string) []byte {
	contentfile, err := os.ReadFile(filman)
	if err != nil {  
		url := "https://learn.zone01oujda.ma/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt"  // url of the standar.txt 
		cmd := exec.Command("curl", "-o", "standard.txt", url) // cammnd for dowload the file 
		_, er := cmd.Output()
		if er != nil {
			fmt.Println("ERROR : Can't open the file")
			os.Exit(1)
		}
		contentfile, _ = os.ReadFile("standard.txt")
	}
	return contentfile
}
