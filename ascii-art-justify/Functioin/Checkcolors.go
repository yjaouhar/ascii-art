package Func

import (
	"fmt"
	"strconv"
	"strings"
)

// verife color code systems
func Checkcolors(namecolor string) ([]uint8, bool) {
	sl := []uint8{}
	iscorect := true // for name of color
	var colo_code string
	namecolor = strings.ToLower(namecolor)
	if strings.HasPrefix(namecolor, "rgb") {
		colo_code = "rgb"
	} else if strings.HasPrefix(namecolor, "hsl") {
		colo_code = "hsl"
	}
	if colo_code != "" { // color code systems have RGB or HSL
		namecolor = strings.TrimPrefix(namecolor, colo_code)
		namecolor = strings.TrimPrefix(namecolor, "(")
		namecolor = strings.TrimSuffix(namecolor, ")")
		slceint := strings.Split(namecolor, ",")
		for i := 0; i < len(slceint); i++ {
			slceint[i] = strings.TrimSuffix(slceint[i], "%")
			index, er := strconv.Atoi(strings.TrimSpace(slceint[i]))
			if er != nil || len(slceint) != 3 {
				iscorect = false
			}
			sl = append(sl, uint8(index))
		}

	} else if strings.HasPrefix(namecolor, "#") { // color code systems have a dicimal code
		count := 0
		namecolor = strings.TrimPrefix(namecolor, "#")
		if len(namecolor) <= 6 {
			count = 6 - len(namecolor)
			fmt.Println(count)
		for i:=0 ;i<count;i++{
			namecolor+="0"
		}
		fmt.Println(namecolor)
		index1, er1 := strconv.ParseUint(namecolor[:2], 16, 8)
		index2, er2 := strconv.ParseUint(namecolor[2:4], 16, 8)
		index3, er3 := strconv.ParseUint(namecolor[4:], 16, 8)

		if er1 != nil || er2 != nil || er3 != nil {
			iscorect = false
		}

		sl = append(sl, uint8(index1), uint8(index2), uint8(index3))
		}else{
			iscorect=false
		}
	
		
	} else { // color code systems have a name
		colorMap := map[string][]string{
			"red":     {"255", "0", "0"},
			"green":   {"0", "255", "0"},
			"blue":    {"0", "0", "255"},
			"yellow":  {"255", "255", "0"},
			"cyan":    {"0", "255", "255"},
			"magenta": {"255", "0", "255"},
			"gray":    {"128", "128", "128"},
			"white":   {"255", "255", "255"},
			"black":   {"0", "0", "0"},
			"orange":  {"255", "165", "2"},
		}
		color, err := colorMap[namecolor]
		fmt.Println(err)
		if !err {
			iscorect = false
		}

		for i := 0; i < len(color); i++ {
			index, _ := strconv.Atoi(color[i])

			sl = append(sl, uint8(index))
		}

	}

	return sl, iscorect
}
