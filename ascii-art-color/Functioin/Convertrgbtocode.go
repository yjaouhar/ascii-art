package Func

import (
	"fmt"
	"image/color"
)

func Convertrgbtocode(c color.RGBA) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c.R, c.G, c.B)
}
