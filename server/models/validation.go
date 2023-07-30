package models

import (
	"strings"
)

func trim(str string) string {
	return strings.Trim(str, " \n")
}
func isEmpty(str string) bool {
	return strings.Trim(str, " \n") == ""
}
