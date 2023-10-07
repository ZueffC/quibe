package app

import (
	"strings"
)

func BBCodeFinder(value interface{}) string {
	if str, ok := value.(string); ok {
		if strings.Contains(str, "\n") {
			str = strings.ReplaceAll(str, "\n", "</ br>")
		}

		return str
	}
	println("Not ok")
	return ""
}
