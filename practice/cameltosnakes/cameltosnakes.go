package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CamelToSnakeCase(s string) string {
	// Return empty string if input is empty
	if s == "" {
		return ""
	}

	// Validate the input
	for i, r := range s {
		// Only letters are allowed
		if !unicode.IsLetter(r) {
			return s
		}

		// No consecutive uppercase letters
		if i > 0 && unicode.IsUpper(r) && unicode.IsUpper(rune(s[i-1])) {
			return s
		}
	}

	// Cannot end with an uppercase letter
	if unicode.IsUpper(rune(s[len(s)-1])) {
		return s
	}

	var result strings.Builder

	// Build the snake_case version
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}

	return result.String()
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
