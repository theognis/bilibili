package tool

import "strings"

func GetExtension(filename string) string {
	position := strings.Index(filename, ".")
	if position == -1 {
		return ""
	}
	position++
	return filename[position:]
}
