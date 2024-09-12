package Func

import (
	"os"
)

// read a banner file
func Readfile(banner string) []byte {
	var contentfile []byte
	banner += ".txt"
	contentfile, err := os.ReadFile(banner)
	//if have error when read  file
	if err != nil || ((banner == "standard.txt" && len(contentfile) != 6623) || (banner == "shadow.txt" && len(contentfile) != 7463) || (banner == "thinkertoy.txt" && len(contentfile) != 5558)) {
		Curl(banner)
		contentfile, _ = os.ReadFile(banner)

	}
	return contentfile
}
