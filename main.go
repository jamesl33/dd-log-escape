package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

// alpha is the Datadog special character alphabet.
//
// https://docs.datadoghq.com/logs/explorer/search_syntax/#escape-special-characters-and-spaces.
var alpha = []string{
	"+",
	"-",
	"=",
	"&&",
	"||",
	">",
	"<",
	"!",
	"(",
	")",
	"{",
	"}",
	"[",
	"]",
	"^",
	`"`,
	"“",
	"”",
	"~",
	"*",
	"?",
	":",
	`\`,
	"#",
	" ",
}

// exit if there was an error.
func exit(err error) {
	if err == nil {
		return
	}

	defer os.Exit(1)

	fmt.Printf("Error: %s\n", err)
}

func main() {
	if len(os.Args) != 2 {
		exit(fmt.Errorf("%s <string>", os.Args[0]))
	}

	fmt.Println(escape(os.Args[1]))
}

// special returns a boolean indicating whether the given string is a special character, and should be escaped.
func special(str string) bool {
	return slices.Contains(alpha, str)
}

// escape the given string.
func escape(str string) string {
	var builder strings.Builder

	var idx int

	for ; idx < len(str); idx++ {
		if special(string(str[idx])) || special(str[idx:min(len(str), idx+2)]) {
			builder.WriteRune('\\')
		}

		builder.WriteByte(str[idx])
	}

	return builder.String()
}
