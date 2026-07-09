package main

import (
	"fmt"
	"strings"
)

func LastWord(s string) string {
	// Split the string into words.
	words := strings.Fields(s)
	if len(words) > 0 {
		// Return the last word plus a newline.
		return words[len(words)-1] + "\n"
	}
	return "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
