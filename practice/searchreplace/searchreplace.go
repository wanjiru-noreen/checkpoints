package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		return 
	}
	// goes through every rune
	fmt.Println(strings.Map(func(r rune) rune {
		//convert string to rune
		if r == []rune(os.Args[2])[0] {
			return []rune(os.Args[3])[0]
		}
		//leave character unchanged
		return r
	}, os.Args[1]))
}
