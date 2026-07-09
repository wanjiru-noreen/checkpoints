package main

import (
	"fmt"
	"unicode"
)

func RepeatAlpha(s string) string {
	res := ""
	for _, r := range s {
		// Check if the current rune is a letter
		if unicode.IsLetter(r) {
			//Convert the letter to lowercase so that
			//'A' and 'a' both have the same alphabet position.
			rep := unicode.ToLower(r) - 'a' + 1
			// Repeat the original character rep times
			for i := 0; i < int(rep); i++ {
				res += string(r)

			}
		} else {
			// If it's not a letter (space, number, punctuation),just append it once.
			res += string(r)
		}
	}
	return res
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
