package stringconcat

import "strings"

func ConcatStrings(stringsToConcat []string) string {
	return strings.Join(stringsToConcat, "")
}