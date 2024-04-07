package strings

import (
	"fmt"
	"strings"
)

const (
	defaultLanguage = "english"
	defaultName     = "World"
)

var greetings = map[string]string{
	"english": "Hello",
	"french":  "Bonjour",
	"spanish": "Hola",
}

// Hello returns a string greeting the person with the provided name in the requested language.
// Name is trimmed of leading and trailing spaces, and an empty name defaults to 'World'.
// Language is case-insensitive. If the requested language is not supported, English is used.
func Hello(name, language string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = defaultName
	}

	language = strings.ToLower(language)
	greeting, found := greetings[language]
	if !found {
		greeting = greetings[defaultLanguage]
	}
	return fmt.Sprintf("%s, %s!", greeting, name)
}
