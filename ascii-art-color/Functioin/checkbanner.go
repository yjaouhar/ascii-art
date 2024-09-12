package Func

// check the banner 
func Checkbanner(banner string) bool {
	return !(banner == "standard" || banner == "shadow" || banner == "thinkertoy")
}
