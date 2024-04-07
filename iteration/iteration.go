package iteration

import "strings"

// Repeat returns a new string made of 'times' copies of string 'characters'
func Repeat(characters string, times int) string {
	return strings.Repeat(characters, times)
}
