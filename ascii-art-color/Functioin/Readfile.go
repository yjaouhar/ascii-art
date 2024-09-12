package Func

import (
	"os"
)

// read a banner file
// read a banner file
func Readfile(banner string) []byte {
	var contentfile []byte
	banner += ".txt"
	contentfile, err := os.ReadFile("banner/" + banner)
	// if have error when read  file
	if err != nil || ((banner == "standard" && len(contentfile) != 6623) || (banner == "shadow" && len(contentfile) != 7463) || (banner == "thinkertoy" && len(contentfile) != 5558)) {
		Curl(banner)
		contentfile, _ = os.ReadFile("banner/" + banner)

	}
	return contentfile
}
