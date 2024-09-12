package Func

import (
	"os"
)

// read a banner file
func Readfile(banner string) []byte {
	var contentfile []byte
	var err error

	readbanner := banner + ".txt"

	contentfile, err = os.ReadFile(readbanner)
	if err != nil || len(contentfile) != 6623 {
		Curl(readbanner)
		contentfile, _ = os.ReadFile(readbanner)
	}
	return contentfile
}
