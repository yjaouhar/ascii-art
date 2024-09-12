package Func

// calcul len of inputtext b ascii-art
func Count(nbword []string, elm string, matrix [][]string, option string) int {
	lent := ""
	count := elm
	if option == "justify" {
		count = ""
		for _, v := range nbword {
			count += v
		}
	}
	for j := 0; j < 1; j++ {
		for k := 0; k < len(count); k++ {
			lent += (matrix[int(count[k]-32)][j]) // print the asci-art
		}
	}

	return len(lent)
}
